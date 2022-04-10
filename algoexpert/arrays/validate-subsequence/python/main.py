def isValidSubsequence(array, sequence):
	s = 0
	for v in array:
		if s == len(sequence):
			break
		if v == sequence[s]:
			s +=1
	return s == len(sequence)