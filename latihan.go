package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, jumArr, jumArr2, x, tempt, hasil int
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		jumArr += A[i]
	}
	for i = 0; i < n; i++{
		if A[i] >= 10{
			x = A[i]
			for x > 0{
				tempt = x%10
				jumArr2 += tempt
				x = x/10
			}
		}else{
			tempt = A[i]
			jumArr2+= tempt
		}
	}
	hasil = jumArr - jumArr2
	fmt.Print(hasil)
}