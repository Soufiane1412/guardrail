# Guardrail Sentinel: Operating Principles

1. **Reliability Beats Features:** A monitoring tool that goes down is useless. We prioritize system uptime and secure data isolation (SOC2 compliance via Postgres RLS) over shipping new UI features. 
2. **Declarative State Management:** Infrastructure and database schemas must be version-controlled and idempotent. No manual SQL or state changes in production environments.
3. **Separation of Concerns:** The Go backend (Core Logic) remains strictly decoupled from the MERN frontend layer. Data-access rules are pushed to the database engine, not the application layer.