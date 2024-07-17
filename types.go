package main

type Registry struct {
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

type RegistryRequest struct {
	Type     int      `json:"type"`
	Store    bool     `json:"store"`
	Registry Registry `json:"registry"`
}
