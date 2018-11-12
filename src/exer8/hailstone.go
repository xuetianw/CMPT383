package exer8

// TODO: your Hailstone, HailstoneSequenceAppend, HailstoneSequenceAllocate functions

func Hailstone(x uint) uint {
	if x%2 == 0 {
		return x / 2
	} else {
		return 3*x + 1
	}
}

func HailstoneSequenceAppend(a uint) []uint {
	list := []uint{}
	list = append(list, a)
	for list[len(list)-1] != 1 {
		list = append(list, Hailstone(list[len(list)-1]))
	}
	return list

}

func countLenth(x uint) uint {
	if x == 1 {
		return 1
	} else {
		return 1 + countLenth(Hailstone(x))
	}
}

func HailstoneSequenceAllocate(a uint) []uint {

	list := make([]uint, countLenth(a))
	list[0] = a
	for i := 1; i < len(list); i++ {
		temp := list[i-1]
		list[i] = Hailstone(temp)
	}
	return list
}
