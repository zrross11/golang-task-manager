package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/zrross11/golang-task-manager/cmd/taskman/cmd"
)

func main() {
	color.Style{color.Green, color.OpBold}.Println(`
________                    __                                         
|        \                  |  \                                        
 \$$$$$$$$______    _______ | $$   __  ______ ____    ______   _______  
   | $$  |      \  /       \| $$  /  \|      \    \  |      \ |       \ 
   | $$   \$$$$$$\|  $$$$$$$| $$_/  $$| $$$$$$\$$$$\  \$$$$$$\| $$$$$$$\
   | $$  /      $$ \$$    \ | $$   $$ | $$ | $$ | $$ /      $$| $$  | $$
   | $$ |  $$$$$$$ _\$$$$$$\| $$$$$$\ | $$ | $$ | $$|  $$$$$$$| $$  | $$
   | $$  \$$    $$|       $$| $$  \$$\| $$ | $$ | $$ \$$    $$| $$  | $$
    \$$   \$$$$$$$ \$$$$$$$  \$$   \$$ \$$  \$$  \$$  \$$$$$$$ \$$   \$$`)
	fmt.Println()

	rootCmd := cmd.NewRootCmd()

	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
