# Inventory-Management-System

📦 Inventory Management System
A web-based inventory management system built with Go (Gin + GORM) and PostgreSQL, allowing users to manage products, stock levels, purchase orders, and sales efficiently.

🚀 Features
🔹 CRUD operations for Products, Inventory, Purchase Orders, and Sales Orders

🔹 PostgreSQL database integration with GORM

🔹 RESTful API with Gin framework

🔹 Auto-timestamping for order dates

🔹 JSON-based request/response structure

🛠️ Tech Stack
Layer	Technology
Language-	Go (Golang)
Web Framework-	Gin
ORM-	GORM
Database-	PostgreSQL
API Format-	JSON (REST)


⚙️ API Endpoints
🔹 Products (/api/products)
GET /: Get all products

GET /:id: Get a product by ID

POST /: Create a new product

PUT /:id: Update product

DELETE /:id: Delete product

🔹 Inventory (/api/inventory)
Same CRUD operations as above

🔹 Purchase Orders (/api/purchaseorder)
POST /: Create purchase order

GET /: List all orders

etc.

🔹 Sales Orders (/api/salesorder)
CRUD for sales transactions

🧪 Sample cURL Requests
# Create a new product
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Mouse","description":"Wireless Mouse","price":500,"quantity":20}'

# Get all products
curl http://localhost:8080/api/products

🧰 Setup Instructions
1. Clone the Repository
bash
Copy code
git clone https://github.com/shubham2216/inventory-management-system.git
cd inventory-management-system


2. Configure PostgreSQL
Update the connection string in your main.go:

go
Copy code
dsn := "host=localhost user=postgres password=yourpassword dbname=db_name port=5432 sslmode=disable"


3. Run the Application
bash

Copy code:-
go mod tidy
go run main.go


📄 License
This project is licensed under the MIT License. See the LICENSE file for details.

🙌 Acknowledgments
Gin Gonic

GORM

PGadmin for managing PostgreSQL visually

