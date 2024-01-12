
# GitHub Copilot

## Quiz Game

This is a simple command-line quiz game written in Go. The game reads a CSV file for a list of questions and answers, starts a timer, and runs through the questions. The user's score is calculated based on the number of correct answers.

### CSV File Format

The CSV file should have two columns: the first column for the questions and the second column for the answers. Here's an example:

> "What is the capital of France?",Paris
> "What is 2+2?",4

### Usage

You can customize the game by passing command-line flags when you start the game:

- **file**: This flag tells the program where the test questions are. The default value is problems.csv.

- **time**: This flag sets the time limit for the test in seconds. The default value is 30.

- **random**: This flag rearranges the set of questions each time the test is taken. The default value is an empty string, which means the questions will not be randomized. To randomize the questions, pass true to this flag.

Here's an example of how to start the game with custom options:

This command starts the game with questions from `my_questions.csv`, a time limit of 60 seconds, and randomizes the order of the questions.

> go run main.go -file=my_questions.csv -time=60 -random=true

### Contribution

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
