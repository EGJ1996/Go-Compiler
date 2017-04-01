package main
import(
    "image/color"
    "math/cmplx"
    "fmt"
    "links"
)
/*func main(){
    const(
        xmin,ymin,xmax,ymax=-2,-2,2,2
        width,height=1024,1024
    )
    img:=image.NewRGBA(image.Rect(0,0,width,height))
    for py:=0;py<=height;py++{
        y:=float64(py)/height*(ymax-ymin)+ymin
        for px:=0;px<=width;px++{
            x:=float64(px)/width*(xmax-xmin)+xmin
            z:=complex(x,y)
            img.Set(px,py,mandelbrot(z))
        }
    }
    png.Encode(os.Stdout,img)
}*/ //code for mandelbrot

func main(){
    var a [3] int=[3]int{3,4,5}
    for i,v:=range a{
        fmt.Printf("%d %d\n",i,v)
    } 
    q:=[...]int{1,2,3}
    fmt.Printf("%T\n",q)
    var arr=make([]int,5)
    one(arr)
    for _,v:=range arr{
        fmt.Printf("%d\n",v)
    }
}
func one(a *[5]int){
    for i:=0;i<len(a);i++{
        a[i]=1;
    }
}
func mandelbrot(z complex128) color.Color{
    const iterations=200
    const contrast=15
    var v complex128
    for n:=uint8(0);n<iterations;n++{
        v=v*v+z
        if cmplx.Abs(v)>2{
            return color.Gray{255-contrast*n}
        }
    }
    return color.Black
}
func appendInt(x []int,int){
    var z []int
    zlen:=len(x)+1
    if(zlen<=cap(x)){
        z=x[:zlen]
    }
    else{
        zcap:=zlen
        if zcap<2*len(x){
            zcap=2*len(x)
        }
        z=make([]int,zlen,zcap)
        copy(z,x)
    }
    z[len(x)]=y
    return z
}
func crawl(url string) []string{
    fmt.Println(ur)
    list,err:=links.Extract(url)
    if err!=nil{
        log.Print(err)
    }
    return list
}

func breadthFirst(f func(item,string) []string,worklist []string){
    seen:=make(map[string]bool)
    for len(worklist)>0{
        items:=worklist
        worklist=nil
        for _,item:=range items{
            if(!seen[item]){
                seen[item]=true
                worklist=append(worklist,f(item)...)
            }
        }
    }
}

var rmdirs []func()

for _,d:=range tempDirs(){
    dir:=d
    os.MkdirAll(dir,0755)
    rmdirs=append(rmdirs,func(){
        os.RemoveAll(dir)
    })
}

for _,rmdir:=range rmdirs{
    rmdir()
}