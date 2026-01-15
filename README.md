# Golang CRUD - Basic Web-Based Management System 🚀

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Version](https://img.shields.io/badge/Version-1.0.0-blue)
![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)

## 📋 Overview
**Golang CRUD** is a streamlined web application designed to handle efficient data management through a robust CRUD (Create, Read, Update, Delete) method. This project focuses on providing a clean user interface while maintaining high data integrity and performance, making it an ideal solution for administrative tasks or inventory tracking.

## ✨ Key Features
* **Full CRUD Functionality:** Seamlessly manage records with optimized database queries.
* **Responsive Dashboard:** A modern, mobile-friendly interface accessible from any device.
* **Data Validation:** Secure server-side and client-side validation to ensure data consistency.
* **Search & Filter:** Quickly locate specific entries with an integrated search mechanism.
* **Dynamic UI Updates:** Real-time feedback for user actions (Success/Error alerts).

## 🛠 Tech Stack
* **Frontend:** HTML5, CSS3 With Bootstrap framework, free and open-source CSS framework directed at responsive, mobile-first front-end web development
* **Backend:** Golang With Echo HTTP Framework, high-performance web framework for building robust and scalable applications in Go
* **Database:** MySQL, open source relational database management system (RDBMS) that’s used to store and manage data. Its reliability, performance, scalability, and manipulate structured data

## 📸 System Preview
| List View | Add New Record | Edit Record |
![Capture](https://github.com/user-attachments/assets/519d6925-1ad5-411b-a6e6-c96ca2ae3ebf)


## 🚀 Getting Started

### Prerequisites
Before you begin, ensure you have the following installed:
* Web Server (XAMPP)
* Database Management System (MySQL)

### Installation
1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/abedehakims/Website-CRUD-MyBini.git](https://github.com/abedehakims/Website-CRUD-MyBini.git)
    ```
2.  **Database Configuration:**
    * Create a new database named `mybini_db`.
    * Import the `database.sql` file (if provided) into your MySQL manager.
    * Update your database credentials in the configuration file (`config.php` or `.env`).
3.  **Deploy:**
    * Move the project folder to your `htdocs` or public directory.
    * Open your browser and navigate to `http://localhost/Website-CRUD-MyBini`.

## 🛡 Security Features
* **SQL Injection Prevention:** Uses prepared statements for all database interactions.
* **XSS Protection:** Sanitized input/output to prevent cross-site scripting.
* **CSRF Protection:** Secure token handling for form submissions.

## 🤝 Contributing
Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---
**Developed with ❤️ by Abede Hakims**
