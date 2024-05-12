package main

func adder(s []int) []int {

	for i := 0; i < len(s); i++ {
		s[i] *= 2
	}
	return s
}

//
//got := adder(value)
//want := {3,4,5,6,7,8,9,10,11}
//if got != want {
//	t.Errorf("Add(1, 2) = %d; want %d", got, want)
//}
