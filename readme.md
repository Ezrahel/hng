Number Facts API

This Go package allows you to fetch number facts and classify a given number based on properties such as being prime, perfect, Armstrong, even, and odd. It integrates with the NumbersAPI service and provides a detailed JSON response with various facts about the number.
Features

    Fetches math-related facts about a given number using the NumbersAPI.
    Classifies numbers based on various properties, including:
        Whether the number is prime.
        Whether the number is perfect.
        Whether the number is an Armstrong number.
        Whether the number is even or odd.
    Calculates the digit sum and returns it along with the facts.
    Returns a structured response in JSON format.

API Endpoint
GET /api/classify-number?number=<number>

This endpoint allows you to classify a number and fetch its facts. It requires a query parameter number, which is the number to classify.
Example Request

GET http://localhost:8080/api/classify-number?number=371

Example Response (JSON)

{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["Armstrong", "odd"],
    "digit_sum": 3,
    "fun_fact": "371 is the sum of cubes of its digits"
}

Functions
NumCheck(ArmstrongNum int) []string

Classifies a number into properties like "Armstrong", "even", "odd" based on the number provided.

    Armstrong Number: A number that is equal to the sum of its own digits raised to the power of the number of digits.
    Even/Odd: Checks whether the number is even or odd.

IsOdd(isOdd int) string

Returns whether a number is "odd" or "even".
IsPerfect(isPerfect int) bool

Checks if the given number is a perfect number. A perfect number is a positive integer that is equal to the sum of its proper divisors, excluding the number itself.
IsPrime(isPrime int) bool

Determines whether a given number is prime. A prime number is a natural number greater than 1 that cannot be formed by multiplying two smaller natural numbers.
ApiCheck(number int) (NumberFactResponse, error)

Fetches a fun fact for the given number from the NumbersAPI service. It returns a NumberFactResponse struct containing the number fact.
NumberHandler(w http.ResponseWriter, r *http.Request)

HTTP handler for the /api/classify-number endpoint. It extracts the number query parameter from the request, calls relevant classification functions, fetches a number fact, and returns a structured JSON response with the classification and number facts.
How to Run the Project
Prerequisites

    Go (1.16+)
    An active internet connection (to fetch number facts from NumbersAPI).

Steps

- Clone the repository:

``` ```bash
      git clone https://github.com/ezrahel/hng.
      cd hng``` ``` 

- Install dependencies:

      go mod tidy

- Run the Go server:

      go run main.go

The server will start listening on port 8000.

- Test the API by visiting:

      http://localhost:8080/api/classify-number?number=371

      You can replace 371 with any other number to classify.

# Example Test

You can test the functionality using curl or Postman to make GET requests.

go
curl "http://localhost:8080/api/classify-number?number=371"


The server will respond with a JSON object containing the classification of the number, such as whether it is prime, perfect, Armstrong, etc.
Code Structure

    NumCheck(): Classifies the number based on Armstrong and even/odd properties.
    IsOdd(): Returns "odd" or "even".
    IsPerfect(): Determines if the number is perfect.
    IsPrime(): Checks if the number is prime.
    ApiCheck(): Fetches a fun fact about the number from NumbersAPI.
    NumberHandler(): HTTP handler to classify numbers and serve the facts as a JSON response.

Notes

    The NumbersAPI URL used here is a public API and may have rate limits. If you're planning to use this in a production environment, be sure to check the NumbersAPI documentation for any usage limitations.
    The code uses Go's built-in http and json packages to handle API requests and responses.

