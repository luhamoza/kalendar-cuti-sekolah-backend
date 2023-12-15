# Holiday API Readme

This is a simple Go (Golang) application that provides an API to retrieve holiday information. The application reads holiday data from a JSON file named `holidays.json` and exposes an endpoint `/holidays` to fetch this information.

## Prerequisites

Before running the application, ensure that you have Go installed on your machine. You can download and install Go from the official [Go website](https://golang.org/dl/).

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/luhamoza/kalendar-cuti-sekolah-backend
   ```

2. Change into the project directory:

   ```bash
   cd kalendar-cuti-sekolah-backend
   ```

3. Build and run the application:

   ```bash
   go mod tidy
   go run main.go
   ```

   This will start the server on `localhost:8000`.

## API Endpoint

- **Endpoint:** `/holidays`
- **Method:** `GET`
- **Description:** Retrieves a list of holiday groups with associated details.

## JSON Data Structure

The holiday data is structured in JSON format and is loaded from a file named `holidays.json`. The JSON file contains an array of holiday groups, where each group has the following structure:

```json
{
  "group": "Group Name",
  "state": ["State 1", "State 2"],
  "session": "Session Name",
  "holidays": [
    {
      "name": "Holiday Name",
      "dateStart": "Start Date",
      "dateEnd": "End Date",
      "totalDays": "Total Days"
    },
    // Additional holidays...
  ]
},
// Additional groups...
```

- `group`: The name of the holiday group.
- `state`: An array of states associated with the group.
- `session`: The session or occasion for the holidays.
- `holidays`: An array of holiday objects, each containing the name, start date, end date, and total days of the holiday.

## Usage

After starting the server, you can access the holiday data by making a GET request to the `/holidays` endpoint. You can use tools like `curl` or Postman, or simply access the endpoint through a web browser.

Example:

```bash
curl http://localhost:8000/holidays
```

## CORS Configuration

The application uses the `gorilla/handlers` package to enable CORS (Cross-Origin Resource Sharing). This allows the API to be accessed from different domains. CORS is configured to allow all origins (`*`), methods, and headers.

## Issues and Contributions

If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request on the GitHub repository.

Enjoy exploring holiday information with this simple Holiday API!
