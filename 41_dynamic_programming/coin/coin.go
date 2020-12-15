package coin

var cho = []int{10, 5, 1}

func CoinTemp(n int) {
	var ret = make(map[int]int)
	for i := 0; i < n; i++ {

	}
}

func choice(res int) int {
	for _, c := range cho {
		if res > c {
			continue
		}
		//res =
		if res == 0 {
			return 1
		}
		if res < 0 {
			return 0
		}
	}
}
