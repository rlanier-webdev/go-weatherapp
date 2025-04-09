### TC001 - Search with Invalid Zip Code

- **Title**: Search with invalid zip code "00000"
- **Precondition***: App is loaded and search input is available
- **Steps**:
  1. Enter "00000" into the search input
  2. Click the "Search" button
- **Expected Result**: A clear error message like "Location not found" is shown
- **Actual Result**: No result or feedback displayed
- **Status**: Fail (see Bug #101)

### TC002 - Invalid Characters in Zip Code Search (Letters)

- **Precondition**: App is loaded, and the zip code search input is available.
- **Test Steps**:
  1. Enter any non-numeric characters (e.g., "ABCDE") into the search input field.
  2. Click the "Search" button.
- **Expected Result**: 
  - The app should display an error message, such as "Please enter a valid zip code" or "Invalid zip code format."
- **Actual Result**: 
  - No result or feedback displayed
- **Status**: Fail (see Bug #102)
