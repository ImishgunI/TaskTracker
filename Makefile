
all: task-cli

task-cli:
	go build -o $@ main.go

clean:
	rm task-cli

rebuild:
	all