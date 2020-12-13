package day13

import (
	"math"
	"math/big"
	"strconv"
	"strings"
	"year2020/helpers"
)

var input, _ = helpers.ProcessInput("day13.input")

func extractData(data []string) (timestamp int, results []int) {
	timestamp, _ = strconv.Atoi(data[0])
	results = make([]int, 0)

	data = strings.Split(data[1], ",")

	for i := 0; i < len(data); i++ {
		if data[i] == "x" {
			results = append(results, 0)
		} else {
			number, _ := strconv.Atoi(data[i])
			results = append(results, number)
		}
	}

	return
}

func findEarliestBus(afterTimestamp int, buses []int) (id, timestamp int) {
	timestamp = -1

	for b := 0; b < len(buses); b++ {
		if buses[b] == 0 {
			continue
		}

		firstTimestamp := int(math.Ceil(float64(afterTimestamp)/float64(buses[b]))) * buses[b]

		if timestamp == -1 || timestamp > firstTimestamp {
			timestamp = firstTimestamp
			id = buses[b]
		}
	}

	return
}

func prepareRemainderData(data []int) [][]*big.Int {
	results := make([][]*big.Int, 2)
	results[0] = make([]*big.Int, 0)
	results[1] = make([]*big.Int, 0)

	for d := 0; d < len(data); d++ {
		if data[d] == 0 {
			continue
		}

		remainder := d
		for {
			if remainder <= 0 {
				break
			}
			remainder = remainder - data[d]
		}

		results[0] = append(results[0], big.NewInt(int64(-remainder)))
		results[1] = append(results[1], big.NewInt(int64(data[d])))
	}

	return results
}

func calculateChineseRemainder(a, n []*big.Int) int {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return int(x.Mod(&x, p).Int64())
}

func SimpleSolution() string {
	baseTimestamp, data := extractData(input)

	id, timestamp := findEarliestBus(baseTimestamp, data)

	return strconv.Itoa((timestamp - baseTimestamp) * id)
}

func AdvancedSolution() string {
	_, data := extractData(input)
	remainderData := prepareRemainderData(data)

	return strconv.Itoa(calculateChineseRemainder(remainderData[0], remainderData[1]))
}
