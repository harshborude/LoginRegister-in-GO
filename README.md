# Go Auth API

A robust, production-ready **Authentication and Role-Based Access Control (RBAC) REST API** built with **Go**.  
This project provides a secure foundation for user management, featuring **dual JWT authentication, password hashing, request validation, and administrative controls**.

---

## 🚀 Features

### User Authentication
Secure **registration and login** using **bcrypt** for password hashing.

### Dual JWT Architecture
- **Short-lived Access Tokens** for seamless API requests.
- **Long-lived Refresh Tokens** for maintaining secure sessions.

### Secure Token Refresh
The refresh endpoint includes:
- **Token rotation** (generating a new refresh token upon use)
- **Account status verification** to ensure disabled accounts cannot obtain new tokens.

### Session Management
Logout functionality that **securely revokes the active refresh token** in the database.

### Role-Based Access Control (RBAC)
Custom **middleware** to restrict endpoints based on user roles:
- `USER`
- `ADMIN`

### Input Validation
Incoming JSON requests are strictly validated using **struct tags**, such as:
- Email formatting
- Password length
- Alphanumeric checks

### Admin Capabilities
Protected administrative routes allow admins to:
- View **all registered users**
- Promote **standard users to admin**

### Database Auto-Migration
Automatically creates and updates database schemas using **GORM**, including a **default seeder for the initial Admin account**.

---

## 🛠️ Tech Stack

| Category | Technology |
|--------|-----------|
| Language | Go (Golang) 1.25+ |
| Framework | Gin (Fast HTTP web framework) |
| Database | PostgreSQL |
| ORM | GORM |
| Security | `golang-jwt/jwt/v5`, `golang.org/x/crypto/bcrypt` |
| Validation | `go-playground/validator/v10` |
| Environment Management | `joho/godotenv` |
