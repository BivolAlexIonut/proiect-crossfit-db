# Go SQL Database Interface

O aplicație de backend dezvoltată în **Go** care servește ca interfață pentru gestionarea unei baze de date **SQL**. Proiectul demonstrează integrarea între un limbaj puternic tastat și un sistem de management al bazelor de date relaționale.

## 🚀 Caracteristici
* **Operații CRUD:** Implementare completă pentru Create, Read, Update și Delete.
* **SQL Integration:** Conexiune securizată și interogări optimizate către baza de date.
* **Go Concurrency:** Utilizarea rutinelor specifice Go pentru eficiență (dacă este cazul).
* **Interfață curată:** Structură modulară a codului pentru o mentenanță ușoară.

## 🛠 Tehnologii folosite
* **Limbaj:** Go (Golang)
* **Bază de date:** SQL (PostgreSQL / MySQL / SQLite)
* **Drivere:** `database/sql` standard library

## 📋 Cum funcționează
1. **Conectare:** Aplicația stabilește o conexiune cu serverul SQL folosind variabile de mediu pentru securitate.
2. **Execuție:** Primește comenzi prin interfață și le traduce în query-uri SQL valide.
3. **Parsare:** Rezultatele din baza de date sunt mapate direct pe structuri (structs) în Go.

## ⚙️ Instalare și Rulare
```bash
# Clonează repository-ul
git clone [https://github.com/utilizator/proiect-db-go.git](https://github.com/utilizator/proiect-db-go.git)

# Instalează dependențele
go mod tidy

# Rulează aplicația
go run main.go
