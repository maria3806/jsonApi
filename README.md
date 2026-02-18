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

