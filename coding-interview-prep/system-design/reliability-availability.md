# System Design: Reliability & Availability

Reliability and availability are crucial aspects of system design, especially for mission-critical applications. They often go hand-in-hand but represent distinct concepts.

## Reliability

**Reliability** is the probability that a system will perform its intended function without failure for a specified period under specified conditions. It's about how often a system fails and how well it recovers from failures.

### Key Concepts for Reliability:

-   **Fault Tolerance:** The ability of a system to continue operating correctly even when some of its components fail.
    -   **Redundancy:** Having duplicate components (hardware, software, data) so that if one fails, another can take over.
        -   **Active-Active:** All redundant components are active and processing requests.
        -   **Active-Passive:** One component is active, and the redundant component is on standby, ready to take over.
    -   **Replication:** Maintaining multiple copies of data or services across different nodes or locations.
-   **Error Handling & Recovery:** Mechanisms to detect, report, and recover from errors gracefully.
    -   **Retries:** Automatically re-attempting failed operations.
    -   **Circuit Breakers:** Preventing a system from repeatedly trying to access a failing service, allowing it to recover.
    -   **Timeouts:** Limiting the duration an operation can take to prevent indefinite blocking.
-   **Monitoring & Alerting:** Continuously observing system health and performance, and notifying operators of potential issues.
-   **Testing:** Thorough testing (unit, integration, system, chaos engineering) to identify and mitigate potential failure points.
-   **Idempotency:** An operation is idempotent if it can be applied multiple times without changing the result beyond the initial application. Important for retries.

## Availability

**Availability** is the proportion of time a system is in a functioning condition. It's often expressed as a percentage (e.g., "five nines" means 99.999% availability).

### Key Concepts for Availability:

-   **Uptime:** The period during which a system is operational.
-   **Downtime:** The period during which a system is not operational.
-   **Mean Time Between Failures (MTBF):** The predicted elapsed time between inherent failures of a system during operation. A higher MTBF indicates greater reliability.
-   **Mean Time To Recover (MTTR):** The average time required to repair a failed component or system. A lower MTTR indicates faster recovery and thus higher availability.
-   **SLA (Service Level Agreement):** A contract between a service provider and a customer that defines the level of service expected.
-   **SLO (Service Level Objective):** A key element within an SLA that defines a measurable characteristic of the service (e.g., availability percentage).
-   **SRE (Site Reliability Engineering):** A discipline that applies software engineering principles to infrastructure and operations problems.

### Achieving High Availability:

-   **Redundancy:** As mentioned in reliability, redundancy is key to availability.
-   **Failover:** Automatically switching to a redundant or standby system upon the failure or abnormal termination of a previously active system.
-   **Disaster Recovery:** A set of policies and procedures to enable the recovery or continuation of vital technology infrastructure and systems following a natural or human-induced disaster.
-   **Geographic Distribution:** Deploying services across multiple data centers or regions to protect against regional outages.
-   **Load Balancing:** Distributes traffic, and can also route around failed servers.

## Go Example: (Conceptual) Health Check

See the conceptual example code in `reliability-availability.go`.
```go
// To run this code:
// 1. Navigate to the system-design directory: cd coding-interview-prep/system-design
// 2. Run the file: go run reliability-availability.go
```
