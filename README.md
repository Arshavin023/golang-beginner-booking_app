# Go Conference Booking App

A simple command-line conference ticket booking application written in **Go**.  
This project demonstrates core Go concepts such as structs, slices, functions, goroutines, wait groups, and basic concurrency control.

---

## 📌 Features

- Book tickets for a conference via CLI input
- Validate user input (name, email, ticket count)
- Track remaining tickets
- Store booking details in memory
- Send booking confirmations asynchronously using goroutines
- Synchronize goroutines using `sync.WaitGroup`
- Designed with concurrency safety in mind using `sync.Mutex`

---

## 🛠️ Technologies Used

- **Go (Golang)**
- Standard library packages:
  - `fmt`
  - `sync`
  - `time`

---

## 🚀 How the Application Works

1. The application greets the user and displays the total and remaining tickets.
2. The user provides:
   - First name
   - Last name
   - Email address
   - Number of tickets to book
3. User input is validated before proceeding.
4. Tickets are booked and stored in memory.
5. A confirmation email is simulated and sent asynchronously.
6. The application waits for all goroutines to complete before exiting.

---

## 🧵 Concurrency Overview

- **Goroutines** are used to send tickets asynchronously.
- **`sync.WaitGroup`** ensures the main program waits for all ticket-sending operations.
- **`sync.Mutex`** is included to protect shared resources:
  - `remainingTickets`
  - `bookings`

> ⚠️ The mutex is currently commented out in `bookTickets`.  
> It should be enabled when ticket booking is handled concurrently.

---

## 📂 Project Structure

```text
.
├── main.go
└── README.md
