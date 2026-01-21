# 🛡️ Guardrail: 90-Day Platform Executive Sprint

> **Goal:** Break into Platform Engineering (Amsterdam) and track toward CTO.
> **Current Status:** Phase 1 - Bootstrap

---

## 📊 High-Level Progress
- [x] **Phase 1: Bootstrap & Foundations** (Days 1-15) - *In Progress*
- [ ] **Phase 2: Cloud & CI/CD** (Days 16-40)
- [ ] **Phase 3: Kubernetes & Guardrail Engine** (Days 41-70)
- [ ] **Phase 4: Leadership & System Design** (Days 71-90)

---

<details open>
<summary><b>🚀 Phase 1: Bootstrap (Days 1-15)</b></summary>

| Day | Title | Description | Status |
|:---:|:---|:---|:---:|
| 1 | The Scaffolding | Set project folder structure, init Git, create GitHub repo. | ✅ |
| 2 | The Desired State | Terraform declarative mindset + write first .tf. | [ ] |
| 3 | State Management | Understand tfstate, backends, sensitive outputs. | [ ] |
| 4 | The Go Provider | Write a Go script that reads a TF-generated file. | [ ] |
| 5 | Clean Code I | Apply naming + small functions to Go code. | [ ] |
| 6 | Variables | Intro to variables, types, validation. | [ ] |
| 7 | Outputs & Locals | Produce structured outputs + locals. | [ ] |
| 8 | Modules I | Create reusable module skeleton. | [ ] |
| 9 | Modules II | Build full VPC module locally. | [ ] |
| 10 | First K8s Deploy | Deploy hello-world app to local cluster. | [ ] |
| 11 | CLI Scaffolding | Start building the Go CLI with Cobra. | [ ] |
| 12 | Commands | Implement init, plan, validate. | [ ] |
| 13 | JSON Parsing | Parse Terraform JSON plan output. | [ ] |
| 14 | Testing | Add unit tests for parsing. | [ ] |
| 15 | Cleanup | Apply Clean Code for function size + clarity. | [ ] |

</details>

<details>
<summary><b>☁️ Phase 2: Cloud & CI/CD (Days 16-40)</b></summary>

| Day | Title | Description | Status |
|:---:|:---|:---|:---:|
| 16 | VPC Theory | Understand AWS VPC design. | [ ] |
| 17 | Terraform VPC | Deploy VPC locally with AWS provider (no apply). | [ ] |
| 18 | IGW + NAT | Add internet gateway & NAT gateway. | [ ] |
| 19 | SG Basics | Define ingress/egress firewalls. | [ ] |
| 20 | Go SDK I | Call AWS API via Go. List EC2 instances. | [ ] |
| 21-24 | Identity | IAM, Roles, Policies, and Go Auth. | [ ] |
| 25-30 | Pipelines | GitHub Actions: Build, Test, TF Plan, Manual Apply. | [ ] |
| 31-40 | Observability | Logging, Monitoring, Networking II, Secrets. | [ ] |

</details>

<details>
<summary><b>☸️ Phase 3: Kubernetes & Guardrail (Days 41-70)</b></summary>

| Day | Title | Description | Status |
|:---:|:---|:---|:---:|
| 41-42 | Containers | Dockerfile best practices & Multi-stage builds. | [ ] |
| 43-50 | K8s Core | Pods, Deployments, Services, Ingress, RBAC. | [ ] |
| 51-54 | Observability | Prometheus, Grafana, and K8s Events. | [ ] |
| 55-60 | **Engine v0.5** | Implement Pod & RBAC inspection via Go `client-go`. | [ ] |
| 61-70 | Advanced K8s | Security, GitOps (Flux), and Cluster Hardening. | [ ] |

</details>

<details>
<summary><b>👔 Phase 4: Leadership & CTO Path (Days 71-90)</b></summary>

| Day | Title | Description | Status |
|:---:|:---|:---|:---:|
| 71-73 | System Design | Architecture for 1M users. | [ ] |
| 74-78 | SRE & Scaling | Incident Mgmt, SLOs, Messaging (Kafka), API Design. | [ ] |
| 79-82 | Documentation | Architecture Reviews & Codebase Docs. | [ ] |
| 83-90 | **Final Push** | Portfolio, Amsterdam Interview Prep, CloudGuard 1.0. | [ ] |

</details>

---

## 📝 Current Daily Notes
- **Day 1:** Resolved Git Authentication loop using PAT.
- **Day 2 (Goal):** Execute `terraform apply` for the first time.