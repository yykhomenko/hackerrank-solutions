package main

func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		if grades[i] > 37 && grades[i]%5 != 0 {
			if 5-int(grades[i]%5) < 3 {
				grades[i] += 5 - grades[i]%5
			}
		}
	}

	return grades
}
