# Issue #14 & #15 — Third service proposal and ADR

## Issue #14: Third service proposal

A short proposal is in **docs/proposals/third-service-notifications.md**. It argues that a dedicated **notifications** service would keep reminder and alert logic separate from scheduling, proposes a service boundary and a minimal proto contract (`Send` RPC), and outlines two integration options (API calls notifications after create vs. event-driven). No full implementation was added; the proposal is enough to evaluate and implement later if desired.

## Issue #15: ADR for one technical decision

**docs/adr/0001-agenda-port-interface.md** documents the decision to introduce the **AgendaPort** interface in the API layer instead of depending on the concrete gRPC client. It covers context (testability, coupling, flexibility), the decision (port + adapter), and consequences (pros, cons, neutral). This matches the refactor done in Issue #7.
