### Bug #101 - No feedback when searching for invalid zip code "00000"

- **Reported On**: April 9, 2025
- **Found In**: Weather Dashboard v1.0
- **Environment**: Chrome 123, macOS 14.2
- **Steps to Reproduce**:
  1. Navigate to the weather dashboard
  2. Enter "00000" into the zip code search input
  3. Click the "Search" button
- **Expected Behavior**: User receives a message indicating that the zip code is invalid or not found
- **Actual Behavior**: Nothing happens; no results or error message is displayed
- **Severity**: Medium
- **Priority**: Medium
- **Suggested Fix**: Display a user-friendly error message for invalid or unresolvable zip codes
- **Status**: Open


### Bug #102 - No error message displayed for invalid characters in zip code field

- **Reported On**: April 9, 2025
- **Found In**: Weather Dashboard v1.0
- **Environment**: Chrome 123, macOS 14.2
- **Steps to Reproduce**:
  1. Navigate to the weather dashboard.
  2. Enter "ABCDE" (or any non-numeric characters) into the zip code search input field.
  3. Click the "Search" button.
- **Expected Behavior**: 
  - An error message should be displayed indicating that the input is invalid, such as "Please enter a valid zip code" or "Invalid zip code format."
- **Actual Behavior**: 
  - No error message is shown, and no action is taken by the app.
- **Severity**: Medium
- **Priority**: High
- **Suggested Fix**: 
  - Implement input validation on the zip code field to reject non-numeric characters and display an appropriate error message.
- **Status**: Open
