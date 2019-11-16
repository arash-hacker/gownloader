package main

import (
	"path"
	_ "io"
	"os"
	"fmt"
	_ "gownloader/gownload"
)

func ConcatParts(){
	os.RemoveAll("./final.mp4")
	for i := 0; i < 4; i++ {
		file,e:=os.OpenFile(fmt.Sprintf("%s/part%d",path.Join(os.Args[0],".."),i ),
			 os.O_RDONLY, 0777)
		if e !=nil{
			panic(e.Error())
		}
		f,_:=file.Stat()
		var b =make([]byte, f.Size())
		file.Read(b)
		
		fmt.Println(b[len(b)-10:],f.Size())
		out,_:=os.OpenFile("./final.mp4", os.O_APPEND| os.O_CREATE| os.O_WRONLY | os.O_RDONLY , 0777)
		bb,e:=out.Write(b)
		fmt.Println(bb,e)
	}
}

func main() {
	//d:=gownload.New("http://s2.mihandownload.com/2011/saeed/1/BruceLeeFighting.rar")
	//d:=gownload.New("http://cdn9.git.ir/1398/02/Linkedin%20Designing%20a%20Book-git.ir/001-creating%20a%20book%20in%20indesign-git.ir.mp4")
	//d.Init(4)
	//fmt.Println(d)	
	//d.StartAll()//wait to finish 
	ConcatParts()

}