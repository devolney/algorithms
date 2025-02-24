# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# RUNNING ALGORITHMS
# ==============================================================================

selection-sort:
	@go run algorithms/4-selection-sort/main.go

bubble-sort:
	@go run algorithms/5-bubble-sort/main.go

insertion-sort:
	@go run algorithms/6-insertion-sort/main.go		

recursion:
	@go run algorithms/7-recursion/main.go	

merge-sort:
	@go run algorithms/8-merge-sort/main.go
