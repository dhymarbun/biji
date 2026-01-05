package main

import "fmt"

func PrintBanner() {
	fmt.Println(Cyan + `
 mmmmmm     mmmmmm      mmmmm   mmmmmm  
 ##""""##   ""##""      """##   ""##""  
 ##    ##     ##           ##     ##    
 #######      ##           ##     ##    
 ##    ##     ##           ##     ##    
 ##mmmm##   mm##mm   #mmmmm##   mm##mm  
 """""""    """"""    """""     """"""` + Reset)
	fmt.Println(Green + "Defensive Web Risk & Traffic Monitor\n" + Reset)
}
