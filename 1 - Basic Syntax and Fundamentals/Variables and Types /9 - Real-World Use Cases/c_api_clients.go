package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Building robust API clients using interfaces and custom types.

// Best Practice: Use interfaces to abstract dependencies, enabling easier testing and flexibility.

// HTTPClient is an interface to abstract the Do method of http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// GeoAPI represents the geolocation data structure returned by the API
type GeoAPI struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

// AddressAPI represents the address data structure returned by the API
type AddressAPI struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     GeoAPI `json:"geo"`
}

// CompanyAPI represents the address data structure returned by the API
type CompanyAPI struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

// UserAPI represents the user data structure returned by the API
type UserAPI struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	UserName   string     `json:"username"`
	Email      string     `json:"email"`
	AddressAPI AddressAPI `json:"address"`
	Phone      string     `json:"phone"`
	Website    string     `json:"website"`
	CompanyAPI CompanyAPI `json:"company"`
}

// APIClient handles communication with the API
type APIClient struct {
	client  HTTPClient
	baseURL string
}

// NewAPIClient creates a new APIClient with the given HTTPClient and base URL
func NewAPIClient(client HTTPClient, baseURL string) *APIClient {
	return &APIClient{
		client:  client,
		baseURL: baseURL,
	}
}

// GetUser fetches a user by ID from the API
func (api *APIClient) GetUser(id int) (*UserAPI, error) {
	url := fmt.Sprintf("%s/users/%d", api.baseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request failed: %w", err)
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user UserAPI
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("decoding response failed: %w", err)
	}
	return &user, nil
}

func main() {
	baseURL := "https://jsonplaceholder.typicode.com"
	httpClient := &http.Client{}

	api := NewAPIClient(httpClient, baseURL)
	user, err := api.GetUser(1)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return
	}
	fmt.Printf("User: %+v\n", user)
}
