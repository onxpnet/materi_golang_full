package main

import (
	"fmt"
	"esdm/sebuah_package"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(fungsiSefolder("selamat pagi"))
	fmt.Println(sebuah_package.SebuahFungsiDiPackage("ada paket"))
	// sebuah_package.sebuahFungsi("fungsi internal") //gagal di import karena private
	fmt.Println(sebuah_package.FungsiLain("beda package, sefolder, beda file"))
}
