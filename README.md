# JobSeekerAPI

🚀 A RESTful API for job seekers and companies to post jobs, apply, and manage applications. Built with Golang.

## API Documentation

The API documentation is generated using Swagger. You can access it by running the server and navigating to:

```
http://localhost:8080/swagger/index.html
```

Make sure you have started the server before accessing the documentation.

## Installation & Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/AyKrimino/JobSeekerAPI.git
   ```
2. Navigate to the project directory:
   ```sh
   cd JobSeekerAPI
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run cmd/main.go
   make run # If you want to use the Makefile
   ```

## Example Requests

### Register a Job Seeker

```json
{
  "email": "JobSeeker@jobseeker.com",
  "password": "abcd1234",
  "role": "JobSeeker",
  "firstName": "job",
  "lastName": "seeker",
  "profileSummary": "ps",
  "skills": ["css", "html", "python"],
  "experience": 0,
  "education": "edu"
}
```

### Register a Company

```json
{
  "email": "company@company.com",
  "password": "dcba4321",
  "role": "company",
  "name": "company",
  "headquarters": "hq",
  "website": "company.com",
  "companySize": "big",
  "industry": "indu"
}
```

## License

MIT
