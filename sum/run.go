package sum

import(
	"os"
	"io"
	"fmt"
	//"log"
)

type float float64

var(
	w, h int
	f [][]float
)

func
scanFloat() (float) {
	var f float

	for {
		_, err := fmt.Scanf("%f", &f)
		if err == io.EOF {
			onEOF()
		}else if err == nil {
			break
		}
	}

	return f
}

func
onEOF(){
	for i:= 0 ; i<h ; i++ {
		for j:=0 ; j<w ; j++ {
			fmt.Printf("%f", f[i][j])
			if(j != w-1){
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
	os.Exit(0)
}

func
Run(args []string){

	h = int(scanFloat())
	w = int(scanFloat())

	//fmt.Printf("%d %d\n", w, h)

	if w<1 || h<1 {
		os.Exit(1)
	}

	f = make([][]float, h)
	for i:=0 ; i<h ; i++ {
		f[i] = make([]float, w)
	}
	
	for{
		for i:=0 ; i<h ; i++ {
			for j:= 0 ; j<w ; j++ {
				f[i][j] += scanFloat()
			}
		}
	}
}
