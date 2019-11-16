package gownload
import (
	"sync"
	"path"
	"io"
	"strconv"
	"errors"
	"net/http"
	"fmt"
	"os"
)

const (
	INIT=iota
	DONE 
	DOWNLOADING
	FAILED
)
type download struct{
	addr string
	size int64
	status int
	ranges []string
	n int64
	sync.WaitGroup
}

func New(s string) *download{

	// var _wg sync.WaitGroup
 return &download{
		addr:s,
		size:0,
		ranges :make([]string, 0),
		n:0,
		// wg:_wg,
	}

}
func (g *download) Init(n int) error {
	
	g.n=int64(n)
	c:=&http.Client{}
	req,e:=http.NewRequest("GET", g.addr, nil)
	
	if e!=nil{
		panic("no internet !")
		return errors.New("404 !")
	}
	rs,e:=c.Do(req)
	if e!=nil{
		panic(e.Error())
		return errors.New("404 !")
	}
	if _, ok := rs.Header["Accept-Ranges"]; !ok {
		return errors.New("accept-Ranges !")
	}
	size, err := strconv.ParseInt(rs.Header["Content-Length"][0], 10, 64)
	if err != nil {
		return err
	}
	fmt.Println(">>>",rs.Header["Content-Length"])//>>> [5479424]
	g.size=size
	gsn:=int64(g.size/int64(n))+1
	for i := 0; i < n; i++ {
		g.ranges=append(g.ranges, fmt.Sprintf("bytes=%d-%d",int64(i)*gsn,int64(i)*gsn+gsn-1) );
	}
	return nil
}
func (g *download) Start(n int) {
	fmt.Println("worder ",n)
	c:=http.Client{}
	req,_:=http.NewRequest("GET", g.addr, nil)
	req.Header.Add("Range", g.ranges[n])
	res,_:=c.Do(req)
	//var ff *os.File
	file,e:=os.OpenFile(fmt.Sprintf("%s/part%d",path.Join(os.Args[0],".."),n ), 
		os.O_APPEND| os.O_CREATE| os.O_WRONLY | os.O_RDONLY, 0777)
	// ff=file
	if e!=nil{
		fmt.Println(" CRAETE ")
		panic(e.Error())
	} 
	io.CopyN(file, res.Body, int64(g.size/g.n)+1) 
	res.Body.Close()
	file.Close()
	g.Done()
}
func (g *download) StartAll() {
	for i := 0; i < len(g.ranges); i++ {
		g.Add(1)
		go g.Start(i)

		// g.Start(i)
	}
	g.Wait() 
	fmt.Println("done printing ..")
}

func (g *download) ConcatParts(){
	os.RemoveAll("./final.mp4")
	out,_:=os.OpenFile("./final.mp4", os.O_APPEND| os.O_CREATE| os.O_WRONLY | os.O_RDONLY , 0777)

	for i := 0; i < int(g.n); i++ {
		file,_:=os.OpenFile(fmt.Sprintf("%s/part%d",path.Join(os.Args[0],".."),i ),
			 os.O_RDONLY, 0777)
		/* if e !=nil{
			file,ee:=os.Create(fmt.Sprintf("%s/part%d",path.Join(os.Args[0],".."),i) )
			if ee!=nil {
				panic(ee.Error())
			}
			//panic(e.Error())
		} */
		f,_:=file.Stat()
		var b =make([]byte, f.Size())
		file.Read(b)
		file.Close()
		fmt.Println(len(b))
		fmt.Println(b[len(b)-10:],f.Size())
		bb,e:=out.Write(b)
		fmt.Println(bb,e)
	}
	out.Close()
}
