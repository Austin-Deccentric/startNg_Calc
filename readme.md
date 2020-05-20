## USAGE
This command-line app is to help interns check if they made it to project stage of not.
To select courses you would have to set some flags.
* `-g or -golang`: Includes golang as one of your course
* `-p or -python`: Includes python as one of your course
* `-f or -frontend`: Includes frontend as one of your course
* `-c or -csharp`: Includes C# as one of your course
By default, `-gen or -general` is set to default (Since the genarl track is for everyone)

## NOTE:
The app fails if you input more fields than what is specified in the statment. This is due to to the fact that `fmt.scanf` stores a value in buffer and this breaks the program.