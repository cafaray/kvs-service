module github.com/cafaray/kvs-service

go 1.19

require (
	github.com/cafaray/internal/data v0.0.0-00010101000000-000000000000 // indirect
	github.com/cafaray/internal/server v0.0.0-00010101000000-000000000000 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.6 // indirect
	golang.org/x/crypto v0.0.0-20220817201139-bc19a97f63c8 // indirect
)

replace github.com/cafaray/internal/server => ./internal/server

replace github.com/cafaray/internal/data => ./internal/data
