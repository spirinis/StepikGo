package pack

import "fmt"

// comment wadaw
func some_go_code() string {
	return "adwawad"
}

func main() {
	// здесь должен быть ваш код
	var N, inp int
	fmt.Scanln(&N)
	slice := make([]int, 0, N)
	for i := 1; i <= N; i++ {
		fmt.Scan(&inp)
		slice = append(slice, inp)
		if inp%2 == 0 {
			fmt.Print(inp, " ")
		}
	}

}
