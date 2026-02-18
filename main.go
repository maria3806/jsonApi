package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Recipe struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Cooked bool   `json:"cooked"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Stats struct {
	Total  int `json:"total"`
	Cooked int `json:"cooked"`
}

type Store interface {
	Add(r Recipe) Recipe
	List() []Recipe
	Get(id int) (Recipe, bool)
}

type MemoryStore struct {
	sync.RWMutex
	recipes []Recipe
	nextID  int
}

func (s *MemoryStore) Add(r Recipe) Recipe {
	s.Lock()
	defer s.Unlock()

	s.nextID++
	r.ID = s.nextID
	s.recipes = append(s.recipes, r)

	return r
}

func (s *MemoryStore) List() []Recipe {
	s.RLock()
	defer s.RUnlock()

	copySlice := make([]Recipe, len(s.recipes))
	copy(copySlice, s.recipes)
	return copySlice
}

func (s *MemoryStore) Get(id int) (Recipe, bool) {
	s.RLock()
	defer s.RUnlock()

	for _, r := range s.recipes {
		if r.ID == id {
			return r, true
		}
	}
	return Recipe{}, false
}

type Handler struct {
	store Store
}

func main() {
	h := &Handler{store: &MemoryStore{}}

	http.HandleFunc("/list", h.listHandler)
	http.HandleFunc("/add", h.addHandler)
	http.HandleFunc("/item/", h.itemHandler)
	http.HandleFunc("/stats", h.statsHandler)

	http.ListenAndServe(":8080", nil)
}

func (h *Handler) listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sendJSON(w, http.StatusOK, h.store.List())
}

func (h *Handler) addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var recipe Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil || recipe.Name == "" {
		sendError(w, "invalid input", http.StatusBadRequest)
		return
	}

	result := h.store.Add(recipe)
	sendJSON(w, http.StatusOK, result)
}

func (h *Handler) itemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/item/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, "invalid id", http.StatusBadRequest)
		return
	}

	recipe, ok := h.store.Get(id)
	if !ok {
		sendError(w, "not found", http.StatusNotFound)
		return
	}

	sendJSON(w, http.StatusOK, recipe)
}

func (h *Handler) statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	list := h.store.List()
	cooked := 0

	for _, r := range list {
		if r.Cooked {
			cooked++
		}
	}

	sendJSON(w, http.StatusOK, Stats{
		Total:  len(list),
		Cooked: cooked,
	})
}

func sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, message string, status int) {
	sendJSON(w, status, ErrorResponse{Error: message})
}
