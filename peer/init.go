package peer

var (
	colors = []int{31, 32, 33, 34, 35, 36, 37}
	using  = make(map[int]*int, len(colors))
)

func init() {
	for _, color := range colors {
		using[color] = nil
	}
}
