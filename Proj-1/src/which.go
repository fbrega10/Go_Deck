package main


import(
	"fmt";
	"path/filepath";
	"os"
)


func main(){
	//defining a new which command line utility
	// by giving as input a name of an env, this function
	//will give back it's value whether it is or not in PATH

	fmt.Printf("Start of the program")
	
	
	arguments:= os.Args
	
	if len(arguments) == 1{
		fmt.Printf("give a valid input argument \n")
		return
	}
	file := arguments[1]

	//get the environment variable PATH -> 
	path := os.Getenv("PATH")
	
	//debugging mode on : 

	fmt.Printf("debugging my fav PATH env var: %s \n", path)

	pathSplit := filepath.SplitList(path)


	for _, directory := range pathSplit{
		fullPath := filepath.Join(directory, file) 
		fileInfo, err := os.Stat(fullPath)
		if err == nil{
			mode := fileInfo.Mode()

			//is it a regular file? 
			if mode.IsRegular(){
				//is it executable? 
				fmt.Println(fullPath)
				return
			}
		}
	}
}