package main

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	for i := int32(0); i < 10000; i++ {
		if v1*i+x1 == v2*i+x2 {
			return "YES"
		}
	}

	return "NO"
}

func kangarooBetter(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 > v2 && (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}

	return "NO"
}
