# 🧿 Tiyago
Tiyago is a powerful and modular backend system designed to handle scalable API services with style and structure. Built with Golang, backed by PostgreSQL, and managed with GORM, Tiyago is here to serve fast, clean, and maintainable code.

## ✨ Features
- ⚙️ Clean layered architecture
- 🔌 RESTful APIs with native Go net/http
- 🧱 Modular: Easy to plug in new features
- 🐘 PostgreSQL via GORM ORM
- 🚀 Auto migration + custom migration command system
- 🐳 Docker support for development

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL

### ⚙️ Development Setup

```
1. Clone the repository
    - git clone https://github.com/agastiya/tiyago.git
    - cd tiyago

2. Download dependencies
    - go mod download

3. Configure environment
    Edit environment/local.yml based on your setup

3. Run the project
    - go run .
```


### 🔃 Database Migrations

```
1. Create new migration file
    - go run . create_users_table

2. The migration file will be generated in: database/migrations

3. Run all migration
    - go run . migrate
```


## 🤝 Contributing

Contributions are welcome! 
Feel free to fork this repo, make changes, and submit a pull request.

## 👤 Author

Made with ❤️ by Agastiya