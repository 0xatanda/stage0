# Stage 0 Profile API (Go/Gin)

A simple RESTful API that returns my profile information and a random cat fact from the Cat Facts API.

## 🚀 Endpoint

`GET /me`

### ✅ Example Response
```json
{
  "fact": "A cat will tremble or shiver when it is in extreme pain.",
  "status": "success",
  "timestamp": "2025-10-17T13:43:51.836648+01:00",
  "user": {
    "email": "atanda0x@gmail.com",
    "name": "Atanda Nafiu",
    "stack": "Go/Gin, Python/Django, JavaScript/Node.js, JAVA/Spring Boot, C#/.NET"
  }
}
