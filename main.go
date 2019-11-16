package main

import (
	
	"fmt"
	"gownloader/gownload"
)



func main() {
	//d:=gownload.New("http://s2.mihandownload.com/2011/saeed/1/BruceLeeFighting.rar")
	d:=gownload.New("http://cdn9.git.ir/1398/02/Linkedin%20Designing%20a%20Book-git.ir/001-creating%20a%20book%20in%20indesign-git.ir.mp4")
	d.Init(4)
	fmt.Println(d)	
	d.StartAll()//wait to finish
	d.ConcatParts()

}