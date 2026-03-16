# Issue #5 — Make failing tests pass

## What was done

Tests were fixed by the prior issues (conflict detection, SlotDuration, blocked slots in availability) and by implementing the skipped TestListReservations and updating mocks for the new availability flow.

## Changes

### reservation_test.go

- **TestCreateReservation_FirstVisitDuration:** Updated mock to expect `ListReservations` with window ending at `base + 60*time.Minute` (FirstVisit duration) instead of 30m, matching the fixed `SlotDuration()`.
- **TestListReservations:** Implemented test: mock `ListReservations` to return one reservation, call `uc.List(ctx, doctorID, from, to)`, assert length 1 and that ID, DoctorID, PatientID, Type match.

### availability_test.go

- **TestGetAvailability_HappyPath:** Added mock for `ListBlockedSlots` returning nil (no blocked slots), since GetAvailability now calls it.
- **TestGetAvailability_WithBlockedSlots:** Replaced comment/placeholder with real mock: `ListBlockedSlots` returns one blocked slot 10:00–11:00; assertion unchanged (no slot may start inside that range).

TestCreateReservation_ConflictDetected, TestCreateReservation_BoundaryConflict, and TestGetAvailability_WithBlockedSlots now pass due to Issue #2 and Issue #11 fixes.
