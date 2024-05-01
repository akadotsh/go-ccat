package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	files := os.Args[1:]
	ccat(files)
}


func ccat(args []string) {
	println("files",args)

	println("files",len(args))
 
	if len(args) >= 2 && args[0] == "head" && strings.HasPrefix(args[1], "-n")  {
		conv, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid argument for -n option:", args[2])
			return
		}
		fmt.Println("conv", conv)

	   fmt.Println("arg",args[3])	

	    file,err:=os.Open(args[3])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

 


	   scanner := bufio.NewScanner(file);

	 	count:= conv
		if(count>0){
		for scanner.Scan(){
				fmt.Println(scanner.Text())
				count--
				if count == 0 {
					break
				}
		}
		} 


	}else if len(args) <= 2 {
	 
	  s:=make([]string,0);

	  for _, arg := range args {
	      file,err:=	 os.Open(arg);

		  	if err != nil {
				fmt.Println(err)
				os.Exit(1)
		    }

		   scanner := bufio.NewScanner(file);
		   for scanner.Scan(){
		    s=append(s,scanner.Text())
		   }

	  }

	for _,concateStr:= range s{
		fmt.Println(concateStr)
	}
     

	}


}