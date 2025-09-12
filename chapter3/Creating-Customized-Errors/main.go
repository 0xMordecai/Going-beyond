package main

func main() {}

// -->> Creating Customized Errors
/*
	The errors.New function is very straightforward.

	-->			err := errors.New("Syntax error in the code")
*/

/*
	The fmt.Errorf function, like many of the functions in the fmt
	function, allows for formatting within the string.

	-->			err := fmt.Errorf("Syntax error in the code at line %d", line)
*/
