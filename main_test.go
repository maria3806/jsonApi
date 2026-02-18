package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Add(r Recipe) Recipe {
	args := m.Called(r)
	return args.Get(0).(Recipe)
}

func (m *MockStore) List() []Recipe {
	args := m.Called()
	return args.Get(0).([]Recipe)
}

func (m *MockStore) Get(id int) (Recipe, bool) {
	args := m.Called(id)
	return args.Get(0).(Recipe), args.Bool(1)
}

func TestAddHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		mockSetup  func(m *MockStore)
		wantStatus int
	}{
		{
			name: "success",
			body: `{"name":"Pasta"}`,
			mockSetup: func(m *MockStore) {
				m.On("Add", mock.Anything).
					Return(Recipe{ID: 1, Name: "Pasta"}).
					Once()
			},
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid input",
			body:       `{"name":""}`,
			mockSetup:  func(m *MockStore) {},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := new(MockStore)
			h := &Handler{store: mockStore}

			tt.mockSetup(mockStore)

			req := httptest.NewRequest("POST", "/add", bytes.NewBufferString(tt.body))
			rr := httptest.NewRecorder()

			h.addHandler(rr, req)

			assert.Equal(t, tt.wantStatus, rr.Code)
			mockStore.AssertExpectations(t)
		})
	}
}

func TestMemoryStore(t *testing.T) {
	s := &MemoryStore{}

	r := s.Add(Recipe{Name: "Salad"})
	assert.Equal(t, 1, r.ID)

	item, ok := s.Get(1)
	assert.True(t, ok)
	assert.Equal(t, "Salad", item.Name)
}

func BenchmarkAdd(b *testing.B) {
	s := &MemoryStore{}
	for i := 0; i < b.N; i++ {
		s.Add(Recipe{Name: "Bench"})
	}
}

func TestListHandler(t *testing.T) {
	mockStore := new(MockStore)
	h := &Handler{store: mockStore}

	mockStore.On("List").
		Return([]Recipe{{ID: 1, Name: "Soup"}}).
		Once()

	req := httptest.NewRequest(http.MethodGet, "/list", nil)
	rr := httptest.NewRecorder()

	h.listHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockStore.AssertExpectations(t)
}

func TestListHandler_MethodNotAllowed(t *testing.T) {
	mockStore := new(MockStore)
	h := &Handler{store: mockStore}

	req := httptest.NewRequest(http.MethodPost, "/list", nil)
	rr := httptest.NewRecorder()

	h.listHandler(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestItemHandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		mockSetup  func(m *MockStore)
		wantStatus int
	}{
		{
			name: "success",
			url:  "/item/1",
			mockSetup: func(m *MockStore) {
				m.On("Get", 1).
					Return(Recipe{ID: 1, Name: "Soup"}, true).
					Once()
			},
			wantStatus: http.StatusOK,
		},
		{
			name:       "invalid id",
			url:        "/item/abc",
			mockSetup:  func(m *MockStore) {},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "not found",
			url:  "/item/99",
			mockSetup: func(m *MockStore) {
				m.On("Get", 99).
					Return(Recipe{}, false).
					Once()
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := new(MockStore)
			h := &Handler{store: mockStore}

			tt.mockSetup(mockStore)

			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rr := httptest.NewRecorder()

			h.itemHandler(rr, req)

			assert.Equal(t, tt.wantStatus, rr.Code)
			mockStore.AssertExpectations(t)
		})
	}
}

func TestStatsHandler(t *testing.T) {
	mockStore := new(MockStore)
	h := &Handler{store: mockStore}

	mockStore.On("List").
		Return([]Recipe{
			{ID: 1, Name: "Soup", Cooked: true},
			{ID: 2, Name: "Salad", Cooked: false},
		}).
		Once()

	req := httptest.NewRequest(http.MethodGet, "/stats", nil)
	rr := httptest.NewRecorder()

	h.statsHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockStore.AssertExpectations(t)
}

func TestSendError(t *testing.T) {
	rr := httptest.NewRecorder()
	sendError(rr, "error", http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestMemoryStore_Full(t *testing.T) {
	s := &MemoryStore{}

	r1 := s.Add(Recipe{Name: "Tea"})
	r2 := s.Add(Recipe{Name: "Coffee", Cooked: true})

	assert.Equal(t, 1, r1.ID)
	assert.Equal(t, 2, r2.ID)

	list := s.List()
	assert.Len(t, list, 2)

	_, ok := s.Get(999)
	assert.False(t, ok)
}

func TestAddHandler_MethodNotAllowed(t *testing.T) {
	mockStore := new(MockStore)
	h := &Handler{store: mockStore}

	req := httptest.NewRequest(http.MethodGet, "/add", nil)
	rr := httptest.NewRecorder()

	h.addHandler(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestItemHandler_MethodNotAllowed(t *testing.T) {
	mockStore := new(MockStore)
	h := &Handler{store: mockStore}

	req := httptest.NewRequest(http.MethodPost, "/item/1", nil)
	rr := httptest.NewRecorder()

	h.itemHandler(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestStatsHandler_MethodNotAllowed(t *testing.T) {
	mockStore := new(MockStore)
	h := &Handler{store: mockStore}

	req := httptest.NewRequest(http.MethodPost, "/stats", nil)
	rr := httptest.NewRecorder()

	h.statsHandler(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestSendJSON(t *testing.T) {
	rr := httptest.NewRecorder()

	sendJSON(rr, http.StatusCreated, Recipe{ID: 1, Name: "Test"})

	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Test")
}
