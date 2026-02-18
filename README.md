# JSON API on Recipes theme
#### This project is a simple RESTful JSON API written in Go for managing recipes. It uses an in-memory store with thread-safe access and follows an interface-based architecture. The application demonstrates clean separation of concerns, dependency injection, unit testing, table-driven tests, mocking with stretchr/testify, benchmark testing, and 80%+ code coverage. The goal of the project is to showcase backend fundamentals, testability, and clean code practices in Go. 👩‍💻✨ 

## Most used Techologies in my project:
> Go library net/http,
> 
> Unit testing
> 
> Table-Driven tests
> 
> Mocking with testify
> 
> Benchmark tests
>
> 

 <img width="200" height="250" alt="image" src="https://github.com/user-attachments/assets/ccffd07a-ffbf-498d-b850-eefdb3ab0597" />

> [!NOTE]
> # What is the Mocking?
> Mocking is a testing technique where real dependencies are replaced with fake objects (mocks).
>For example, instead of using a real database, we create a mock store that pretends to save or return data.
> ## Why is this useful?
> * Tests run faster
> * No need for real external systems
> * We can simulate errors easily
> * We test only the logic we care about
> ### In this project, mocking is implemented using the library:
> - stretchr/testify
> ### It allows defining expected method calls and return values.

<img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/67b4d5e9-09c8-45cf-92c1-b2cd44848bf6" />

> [!NOTE]
> # What is Unit Testing?
> Unit testing means testing small parts (units) of the application separately.
> ### Examples of units:
> * One function
> * One handler
> * One store method 
> ## The goal is to check:
> * Correct behavior
> * Edge cases
> * Error handling
> * Invalid input handling
> ### Unit tests improve confidence in code changes and reduce bugs.

<img width="90" height="90" alt="image" src="https://github.com/user-attachments/assets/50dabd72-6f6b-4ed3-8164-91334fb3b288" />

> [!NOTE]
> # What is Benchmarking?
> ### Brenchmaking measures how fast a functions runs.
> ### In Go, benchmark tests use:
> ```
> func BenchmarkSomething(b *testing.B)
> ```
> ### The test runs the function many times and measures performance. 
> ### This helps understand:
> * How efficient the code is
> * Whether performance degrades with scale
> * If optimization is needed
>   ### Even through this project uses an in-memory store, benchmarking demonstrates awareness of performance considerations.
>   
 <img width="400" height="400" alt="image" src="https://github.com/user-attachments/assets/c990d46e-9f1a-4e59-b706-264040b07bd5" />


# How to work with this Project?
### (Includes how to check code coverege %)
## First of all - Run the server
1. Make sure you have Go installed.
 - In the project folder, run:
```
go run main.go
```

The server will start on:
```
http://localhost:8080
```

> [!IMPORTANT]
> ## Available Endpoints
> 
> GET /list — get all recipes
> 
> POST /add — add a new recipe
> 
> GET /item/{id} — get recipe by ID
> 
> GET /stats — get statistics

## You can test endpoints using:

> Browser (for GET requests)
> Postman
> Insomnia (i use this)
> curl

# Example request:

```
POST /add
{
  "name": "Pasta",
  "cooked": true
}
```

# 🧪 How to Run Tests

## To run all tests:
```
go test ./...
```

# 📊 How to Check Code Coverage ‼️

## To generate coverage report:
```
go test -coverprofile=coverage.out ./...
```

## To see coverage percentage in terminal:
```
go test -cover
```

## To open detailed HTML report:
```
go tool cover -html=coverage.out
```

> ![IMPORTANT]
> ## This will open a browser page showing:
> ### Green lines → covered code
> ### Red lines → uncovered code

<img width="200" height="200" alt="image" src="https://github.com/user-attachments/assets/d06a3e34-8213-4635-ac68-75079179a089" />


# Project Reflection
#### A major part of this project was focused on increasing code coverage. I spent four days refining tests, adding edge cases, covering error branches, and improving weak spots in the logic. Reaching 90% coverage was not immediate — it required patience, debugging, and rethinking how I structure tests and mocks.

#### By the way, while the code coverage percentage was increasing, the number of my nervous cells was rapidly decreasing. Still, this experience helped me better understand code quality, reliability, and the importance of writing well-tested software.

<img width="300" height="300" alt="image" src="https://github.com/user-attachments/assets/dae5b2b4-d9c0-4085-82d1-7f9f916afbeb" />



