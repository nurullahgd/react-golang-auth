# React & Go Authentication Project

A full-stack authentication system built with **React** and **Go (Golang)**. This project supports secure user registration, login, logout, and JWT-based session handling.

---

## 🌟 Features

- ✅ User registration & login
- 🔐 JWT-based authentication
- 🛡️ Protected frontend routes
- 🔑 Password hashing with bcrypt
- 🗄️ PostgreSQL database with GORM
- 🌐 CORS configuration support

---

## 🛠️ Tech Stack

### Frontend
- React (with TypeScript)
- React Router DOM
- Bootstrap

### Backend
- Go (Golang)
- Fiber (web framework)
- GORM (ORM for PostgreSQL)
- PostgreSQL
- JWT
- Bcrypt

---

## ⚙️ Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/react-golang-auth.git
```

### 2. Start PostgreSQL with Docker

```bash
cd golang-auth
docker-compose up -d
```

### 3. Run the Backend

```bash
cd golang-auth
go mod download
go run main.go
```

### 4. Run the Frontend

```bash
cd react-auth
npm install
npm start
```

## 🔗 URLs
Frontend: http://localhost:3000

Backend: http://localhost:8000

## 📁 Folder Structure
```bash
react-golang-auth/
│
├── golang-auth/       # Go backend (Fiber, GORM, JWT)
├── react-auth/        # React frontend (TypeScript, Bootstrap)
└── README.md
```
## 🧾 License
This project is open-source and available under the MIT License.
