🛡️ Project: Guardrail Sentinel (90-Day Sprint)

    Goal: Build a production-ready Compliance & Reputation SaaS in Go/MERN. Status: 🟢 Phase 1: The Foundation | Progress: [▓░░░░░░░░░░░░░░░░░░░░░░░░░░░░░] 3%

📅 Phase 1: The Data Fortress (Days 1-30)

Focus: PostgreSQL, Multi-tenancy, and Go Core

    [x] Day 1-3: The Multi-tenant Blueprint

        [x] Design DB Schema for "Companies" vs "Scans".

        [x] Implement Postgres Row-Level Security (RLS).

        [ ] Director Insight: Read "Designing Data-Intensive Apps" (Ch. 2: Data Models).

    [ ] Day 4-10: The Go Engine Start

        [ ] Initialize Go module and folder structure (/cmd, /internal, /pkg).

        [ ] Connect Go to Postgres using pgx (Avoid heavy ORMs for performance).

    [ ] Day 11-20: The "Clean" API

        [ ] Build REST endpoints for "Register Brand" and "Fetch Health".

        [ ] Apply "Clean Code" (Small functions, meaningful names).

    [ ] Day 21-30: Security & Middleware

        [ ] Implement JWT Authentication.

        [ ] Build a "Reputation-based" rate limiter in Go.

📅 Phase 2: The Concurrency Engine (Days 31-60)

Focus: Advanced Go, Goroutines, and System Reliability

    [ ] Day 31-40: The Parallel Scanner

        [ ] Build a worker pool in Go to scan 50+ URLs simultaneously.

        [ ] Director Insight: Master sync.WaitGroup and channels.

    [ ] Day 41-50: Resilience Patterns

        [ ] Implement the "Circuit Breaker" pattern (If a scan fails, don't crash the system).

        [ ] Build advanced error logging (SQL-based audit trail).

    [ ] Day 51-60: The PR Logic

        [ ] Create the "Sentiment Algorithm" (Basic logic: Uptime + Security Headers = Score).

        [ ] Unit test the scoring engine (Aim for 80% coverage).

📅 Phase 3: The Executive Interface (Days 61-90)

Focus: React, Branding, and Global Launch

    [ ] Day 61-75: The "PR-Ready" Dashboard

        [ ] Setup React + Tailwind CSS.

        [ ] Build "Executive Cards" (Big numbers, green/red status lights).

    [ ] Day 76-85: Real-time Sentinel

        [ ] Integrate WebSockets for live status updates.

        [ ] Build the "Crisis Mode" UI (What happens when the brand score drops).

    [ ] Day 86-90: The Sell

        [ ] Deploy to AWS/Postgres.

        [ ] Write the README as a Technical Case Study for your Director application.