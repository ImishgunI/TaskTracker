
all: clean task-cli

task-cli:
	go build -o $@ main.go

clean:
	rm -f task-cli

rebuild:
	all