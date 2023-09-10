PORT_1 := 8080
PORT_2 := 8081
PORT_3 := 8082

SLOW := 6


API_ENDPOINT_1 := http://localhost:$(PORT_1)
API_ENDPOINT_2 := http://localhost:$(PORT_2)
API_ENDPOINT_3 := http://localhost:$(PORT_3)

run-api-1:
	@echo "Running API 1 on port $(PORT_1)"
	PORT=$(PORT_1) go run main.go

run-api-2:
	@echo "Running API 2 on port $(PORT_2)"
	PORT=$(PORT_2) go run main.go

run-api-3:
	@echo "Running API 3 on port $(PORT_3)"
	PORT=$(PORT_3) go run main.go

run-slow-api-1:
	@echo "Running API 1 on port $(PORT_1)"
	PORT=$(PORT_1) slow=$(SLOW) go run main.go 

run-slow-api-2:
	@echo "Running API 2 on port $(PORT_2)"
	PORT=$(PORT_2) slow=$(SLOW) go run main.go

run-slow-api-3:
	@echo "Running API 3 on port $(PORT_3)"
	PORT=$(PORT_3) slow=$(SLOW) go run main.go

