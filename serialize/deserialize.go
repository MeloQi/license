package serialize

import "errors"

func Deserialize(data string) (string, error) {
	if len(data) < 192 {
		return "", errors.New("serialize err, length err")
	}
	l1 := data[:64]
	l2 := data[64:128]
	l3 := data[128:192]
	l4 := data[192:]

	//1. 行行交换
	l1Arry := []byte(l1)
	l2Arry := []byte(l2)
	l3Arry := []byte(l3)
	for i := 0; i < len(l1Arry); i++ {
		switch i % 3 {
		case 0:
			break
		case 1:
			t := l1Arry[i]
			l1Arry[i] = l3Arry[i]
			l3Arry[i] = l2Arry[i]
			l2Arry[i] = t
			break
		case 2:
			t := l1Arry[i]
			l1Arry[i] = l2Arry[i]
			l2Arry[i] = l3Arry[i]
			l3Arry[i] = t
			break
		}
	}
	l1 = string(l1Arry)
	l2 = string(l2Arry)
	l3 = string(l3Arry)

	//2. 首尾交换
	l1 = headtailSW(l1)
	l2 = headtailSW(l2)
	l3 = headtailSW(l3)

	//3. 相邻交换
	l1 = adjacentSW(l1)
	l2 = adjacentSW(l2)
	l3 = adjacentSW(l3)

	//4. l1与l3交换
	lt := l1
	l1 = l3
	l3 = lt

	return l1 + l2 + l3 + l4, nil
}
