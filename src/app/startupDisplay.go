package main

import (
	"fmt"
)

func startupDisplay(port string) {
	fmt.Println(` ___   ___  _______   _______`)
	fmt.Println(` \  \ /  / |   _   | |   ____|`)
	fmt.Println(`  \  V  /  |  |_|  | |  |____`)
	fmt.Println(`  /  _  \  |     __| |   ____|`)
	fmt.Println(` /  / \  \ |  |\  \  |  |`)
	fmt.Println(`/__/   \__\|__| \__\ |__|`)
	fmt.Println(`______________________________`)
	fmt.Println("\nWelcome to the XRF File Tools Program\n")
	fmt.Println("Please access the program via the interface, at:")
	fmt.Println("http://localhost" + port)
	fmt.Println("\nType 'Ctrl+C' or close this window to terminate the program.\n")
	return
}
