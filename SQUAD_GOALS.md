# 🛡️ Project: Guardrail Sentinel (90-Day Sprint)
> **Goal:** Build a production-ready Compliance & Reputation SaaS in Go/MERN.
> **Status:** 🟢 Phase 1: The Foundation | **Progress:** [░░░░░░░░░░] 5%
> **Director's Note:** This system prioritizes Observability, Reliability, and Isolation.

## 📅 Phase 1: The Data Fortress (Days 1-30)
*Focus: PostgreSQL, Multi-tenancy, and Observability*

- [x] **Day 1-3: The Multi-tenant Blueprint** (#1)
    - [x] Design DB Schema for "Companies" vs "Scans" (#2)
    - [x] Implement Postgres Row-Level Security (RLS) (#3)
    - [ ] **[DIRECTOR]** Create `OPERATING_PRINCIPLES.md` (Define "Reliability beats Features").
- [ ] **Day 4-10: The Go Engine & Migrations**
    - [ ] Initialize Go module and pgx connection pool.
    - [ ] **[DIRECTOR]** Setup Database Migrations (No manual SQL in production).
    - [ ] **[DIRECTOR]** Implement Structured Logging (`slog`) with `request_id` and `tenant_id`.
- [ ] **Day 11-20: The "Observable" API**
    - [ ] Build REST endpoints for "Register Brand" and "Fetch Health".
    - [ ] Add basic metrics middleware (Latency, Error Rate tracking).
- [ ] **Day 21-30: Security & Policy**
    - [ ] Implement JWT Authentication.
    - [ ] **[DIRECTOR]** Policy-Driven Throttling (Rate limits based on subscription tier, not just "vibes").

---

## 📅 Phase 2: The Concurrency & Cost Engine (Days 31-60)
*Focus: Advanced Go, Backpressure, and FinOps Thinking*

- [ ] **Day 31-40: The Parallel Scanner**
    - [ ] Build a worker pool to scan 50+ URLs simultaneously.
    - [ ] **[DIRECTOR]** Implement Backpressure (Stop accepting jobs if DB latency > 500ms).
- [ ] **Day 41-50: Resilience & Cost**
    - [ ] Implement Circuit Breakers (Fail fast if a service is down).
    - [ ] **[DIRECTOR]** The Scale Simulation: Document "Cost at 5,000 tenants" in README.
- [ ] **Day 51-60: The Scoring Logic**
    - [ ] Create the "Sentiment Algorithm" (Uptime + Security Headers).
    - [ ] Unit test the scoring engine (80% coverage).

---

## 📅 Phase 3: The Executive Interface (Days 61-90)
*Focus: React, Reporting, and Accountability*

- [ ] **Day 61-75: The Dashboard**
    - [ ] Setup React + Tailwind CSS.
    - [ ] Build "Executive Cards" (Traffic Light Status).
- [ ] **Day 76-85: The "Artifact"**
    - [ ] **[DIRECTOR]** Build "Weekly Report Generator" (PDF/Email summary of risks).
    - [ ] Integrate WebSockets for live alerts.
- [ ] **Day 86-90: The Launch**
    - [ ] Deploy to AWS/Postgres.
    - [ ] Write the Technical Case Study (Focus on: "How I ensured this is profitable and safe").