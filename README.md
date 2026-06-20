# Go Conference Booking App

A command-line conference ticket booking application written in **Go**, built while learning core language and concurrency concepts.

The app simulates a real ticket-booking flow: it collects attendee details, validates them, tracks remaining capacity, and sends a simulated confirmation "email" in the background — all while staying safe under concurrent access.

---

## Features

- Interactive CLI prompts for first name, last name, email, and ticket count
- Input validation:
  - Names must be at least 2 characters and contain letters only
  - Email must contain an `@` symbol
  - Ticket count must be greater than zero and not exceed remaining capacity
- Live tracking of remaining ticket inventory
- In-memory storage of all bookings
- Asynchronous "confirmation email" sending via goroutines
- Graceful shutdown — the app waits for all in-flight confirmations to finish before exiting
- Thread-safe access to shared state (`remainingTickets`, `bookings`) via `sync.Mutex`
- Automatically stops prompting once the conference is sold out

---

## Technologies Used

- **Go** (standard library only — no external dependencies)
  - `fmt` — CLI output and input scanning
  - `sync` — `WaitGroup` and `Mutex` for concurrency control
  - `time` — simulated delay for the confirmation "email"

---

## How It Works

1. The app greets the user and shows total vs. remaining tickets.
2. While tickets remain, it loops:
   - Prompts for first name, last name, email, and ticket count
   - Validates all four fields
   - If valid: books the tickets, prints a confirmation, and kicks off a background goroutine to "send" the ticket
   - If invalid: prints exactly which field(s) failed, and prompts again
3. Once `remainingTickets` reaches 0, the loop exits and the app prints a sold-out message.
4. Before exiting, the app waits (`wg.Wait()`) for every pending "send confirmation" goroutine to finish — so the program won't quit mid-send.

---

## Concurrency Design

This project is a small, deliberate exercise in safe concurrency, not just goroutines for their own sake:

| Mechanism | Purpose |
|---|---|
| `go sendTickets(...)` | Simulates sending a confirmation without blocking the next prompt |
| `sync.WaitGroup` | Ensures `main()` doesn't exit while a confirmation is still "in flight" |
| `sync.Mutex` | Protects `remainingTickets` and `bookings` from concurrent read/write |
| `defer wg.Done()` | Guarantees the wait group is released even if `sendTickets` panics |

The mutex matters even though bookings currently happen one at a time on the main goroutine: `bookTickets` (writer) and `getFirstNames` (reader) both touch shared package-level state, and locking both sides now means the code stays race-safe if this is ever extended to handle multiple bookings concurrently — without that, it'd be a silent landmine waiting for a future feature.

---

## Known Limitations (by design — this is a learning project)

- Data is in-memory only; nothing persists across runs
- No real email sending — `sendTickets` just sleeps and prints
- No CLI flags or config; ticket cap (`conferenceTickets`) is a hardcoded constant
- Single-process only — not built for distributed/horizontal scaling

---

## Project Structure

```text
.
├── main.go         # Entry point, booking flow, concurrency orchestration
├── helper.go        # Input validation logic
└── README.md
```

---

## Running It

```bash
go run .
```

(or explicitly: `go run main.go helper.go`)

Run with the race detector while developing, to confirm there's no data race as the code evolves:

```bash
go run -race .
```

---

## What I Learned

- Structuring a CLI program around structs and slices
- Validating user input cleanly with multiple return values
- Using goroutines + `WaitGroup` for non-blocking background work
- Why `sync.Mutex` matters even in code that *looks* single-threaded at first glance
- Avoiding subtle bugs like unsigned integer underflow (`uint` subtraction going negative) when validating before mutating shared state