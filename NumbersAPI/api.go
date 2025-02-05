package numbersapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"

	// "math"
	"net/http"
	"strconv"
)

type NumberFactResponse struct {
	Numbers    int      `json:"number"`
	IsPrimes   bool     `json:"is_prime"`
	IsPerfects bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSums  int      `json:"digit_sum"`
	// IsOdds     string   `json:"is_odd"`
	FunFacts string `json:"fun_fact"`
}

func NumCheck(ArmstrongNum int) []string {
	prop := []string{}
	total := 0
	numStr := strconv.Itoa(ArmstrongNum)
	totalLen := len(numStr)
	for _, numSt := range numStr {
		toDigit, _ := strconv.Atoi(string(numSt))
		total += int(math.Pow(float64(toDigit), float64(totalLen)))
	}
	if total == ArmstrongNum {
		prop = append(prop, "armstrong")
	}
	if ArmstrongNum%2 == 0 {
		prop = append(prop, "even")
	}
	if ArmstrongNum%2 == 1 {
		prop = append(prop, "odd")
	}

	return prop

}

func IsOdd(isOdd int) string {
	if isOdd%2 == 1 {
		return "odd"
	}
	return "even"
}
func IsPerfect(isPerfect int) bool {
	if isPerfect < 1 {
		return isPerfect < 1
	}
	sumPerfect := 0
	for i := 1; i <= isPerfect/2; i++ {
		if isPerfect%2 == 0 {
			sumPerfect += i
		}
	}
	return sumPerfect == isPerfect
}
func IsPrime(isPrime int) bool {
	if isPrime <= 1 {
		return false
	}
	if isPrime == 2 {
		return true
	}
	if isPrime%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(isPrime))); i += 2 {
		if isPrime%i == 0 {
			return false
		}
	}
	return true
}
func ApiCheck(number int) (NumberFactResponse, error) {
	api := fmt.Sprintf("http://numbersapi.com/%v/math", number)
	getApi, err := http.Get(api)
	if err != nil {
		log.Fatalf("Error fetching api: %v", err)
	}
	defer getApi.Body.Close()

	GetApi, err := io.ReadAll(getApi.Body)
	if err != nil {
		log.Fatalf("Error occuered while decoding response body: %v", err)
	}
	var factResponse NumberFactResponse
	err = json.Unmarshal(GetApi, &factResponse)
	if err != nil {
		return NumberFactResponse{}, err
	}
	return factResponse, nil
}

func NumberHandler(w http.ResponseWriter, r *http.Request) {
	numberInput := r.URL.Query().Get("number")
	if numberInput == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := map[string]interface{}{
			"number": numberInput,
			"error":  true,
		}
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			http.Error(w, "Invalid number input", http.StatusInternalServerError)
		}
		return
	}
	number, err := strconv.Atoi(numberInput)
	totalLen := len(numberInput)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := map[string]interface{}{
			"number": numberInput,
			"error":  true,
		}
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			http.Error(w, "Invalid number input", http.StatusInternalServerError)
		}
		return
	}
	api := fmt.Sprintf("http://numbersapi.com/%v/math", number)
	getApi, err := http.Get(api)
	if err != nil {
		log.Fatalf("Error fetching api: %v", err)
	}
	defer getApi.Body.Close()

	GetApi, err := io.ReadAll(getApi.Body)
	if err != nil {
		log.Fatalf("Error occuered while decoding response body: %v", err)
	}
	isNumber := number
	isPrime := IsPrime(number)
	// isOdd := IsOdd(number)
	numCheck := NumCheck(number)
	// prop = append(prop, isOdd)
	// prop = append(prop, numCheck)
	isPerfect := IsPerfect(number)
	newFactResponse := &NumberFactResponse{
		Numbers:    isNumber,
		IsPerfects: isPerfect,
		IsPrimes:   isPrime,
		// IsOdds:     isOdd,
		Properties: numCheck,
		DigitSums:  totalLen,
		FunFacts:   string(GetApi),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newFactResponse)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
