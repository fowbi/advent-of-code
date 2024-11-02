# See https://gist.github.com/prwhite/8168133
help: ## show help message
	@ awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target> \033[36m\033[0m\nTargets:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

setup-today: ## set up AOC challenge for the current day
	@ go run cmd/setup/main.go

setup: ## set up AOC challenge for given day, day is mandatory
	@ go run cmd/setup/main.go -day $(day)

fetch-input: ## fetch AOC challenge input for today, day is mandatory
	@ go run cmd/input/main.go -day $(day)

fetch-input-today: ## fetch AOC challenge input for given day, day is mandatory
	@ go run cmd/input/main.go

test: ## test AOC challenge for a given day, day is mandatory
	@ go test "./star_$(shell printf "%02d" $(day))" -v
