# Weather Service

This Weather Service is a simple HTTP server written in Go that uses the OpenWeather API to fetch current weather data based on latitude and longitude coordinates.

## Getting Started

These instructions will show you how to get a copy of the project up and running on your local machine and test the api.

### Prerequisites

- Install Go
- Create an account with Openweather at (https://home.openweathermap.org/)
- Get your API key from (https://openweathermap.org/api)

### How to run the server

1. Clone the repository to your local machine:

```sh
git clone https://github.com/ceriustech/weather-service.git
cd weather-service
```

2. You can either create a .env file in the root of the project or add your OpenWeather API key directly in the server code:

- If you choose to create a .env file

```sh
OPENWEATHER_API_KEY=your_api_key_here
apiKey := os.Getenv("OPENWEATHER_API_KEY")
```

- If you choose to insert the key directly into the code

```sh
apiKey := "your_api_key_here"
```

3. Install the necessary Go dependencies:

```sh
go mod tidy
```

4. Run the server

- Navigate to the directory you cloned the project to and use the following command:

```sh
go run main.go
```

### How to test the API

After the server is running, you can fetch weather data by making a GET request to /weather endpoint by setting the lat (latitude) and lon (longitude) query parameters.

(Note\*\* fill in the lat and long with your own coordinates)

1. Go to http://localhost:8080/ in your browser and paste the following

```sh
http://localhost:8080/weather?lat=40.7128&lon=-74.0060
```

2. Use a curl command

```sh
curl "http://localhost:8080/weather?lat=40.7128&lon=-74.0060"
```

3. Use a tool like Postman

- Either download or use it on the web and make a `GET` request with the following endpoint:

```sh
http://localhost:8080/weather?lat=40.7128&lon=-74.0060
```

