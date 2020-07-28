package main

func gradingStudents(grades []int32) []int32 {
	for i := 0; i < len(grades); i++ {
		if grades[i] >= 38 {
			if grades[i]%5 > 2 {
				grades[i] += 5 - grades[i]%5
			}
		}
	}

	return grades
}
