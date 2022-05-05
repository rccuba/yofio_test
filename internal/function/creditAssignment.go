package creditAssignment

import (
	"errors"
	"math/rand"
	"strconv"
	"test_robert_yofio/internal/static"
	"time"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type creditAssigner struct {
}

func NewCreditAssigner() CreditAssigner {
	return &creditAssigner{}
}

func (c *creditAssigner) Assign(investment int32) (int32, int32, int32, error) {
	combinations := make([][]int32, 0)
	for i := 1; i <= int(investment/300); i++ {
		if int32(300*i) == investment {
			combinations = append(combinations, []int32{int32(i), 0, 0})
		}
	}
	for i := 1; i <= int(investment/500); i++ {
		if int32(500*i) == investment {
			combinations = append(combinations, []int32{0, int32(i), 0})
		}
	}
	for i := 1; i <= int(investment/700); i++ {
		if int32(700*i) == investment {
			combinations = append(combinations, []int32{0, 0, int32(i)})
		}
	}
	for i := 1; i <= int(investment/500); i++ {
		for j := 1; j <= int(investment/500); j++ {
			if int32(300*i)+int32(500*j) == investment {
				combinations = append(combinations, []int32{int32(i), int32(j), 0})
			}
		}
	}
	for i := 1; i <= int(investment/700); i++ {
		for j := 1; j <= int(investment/700); j++ {
			if int32(500*i)+int32(700*j) == investment {
				combinations = append(combinations, []int32{0, int32(i), int32(j)})
			}
		}
	}
	for i := 1; i <= int(investment/700); i++ {
		for j := 1; j <= int(investment/700); j++ {
			if int32(300*i)+int32(700*j) == investment {
				combinations = append(combinations, []int32{int32(i), 0, int32(j)})
			}
		}
	}
	for i := 1; i <= int(investment/700); i++ {
		for j := 1; j <= int(investment/700); j++ {
			for k := 1; k <= int(investment/700); k++ {
				if int32(300*i)+int32(500*j)+int32(700*k) == investment {
					combinations = append(combinations, []int32{int32(i), int32(j), int32(k)})
				}
			}
		}
	}
	if len(combinations) > 0 {
		rand.Seed(time.Now().Unix())
		creditAssigner := combinations[rand.Intn(len(combinations))]
		return creditAssigner[0], creditAssigner[1], creditAssigner[2], nil
	}
	return 0, 0, 0, errors.New(static.MsgUndeliveredAmount + " : " + strconv.Itoa(int(investment)))
}
