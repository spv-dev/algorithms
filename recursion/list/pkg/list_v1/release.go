package list

import "errors"

func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + Sum(list[1:])
}

func Length(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return 1 + Length(list[1:])
}

func Max(list []int) (int, error) {
	if len(list) == 0 {
		return 0, errors.New("List is empty")
	}

	if len(list) == 1 {
		return list[0], nil
	}

	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0], nil
		} else {
			return list[1], nil
		}
	}

	max, err := Max(list[1:])

	if err != nil {
		return 0, err
	}

	if max > list[0] {
		return max, nil
	} else {
		return list[0], nil
	}
}
