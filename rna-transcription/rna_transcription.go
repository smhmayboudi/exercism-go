package strand

func ToRNA(dna string) string {
	rna := []rune(dna)
	for i := 0; i < len(rna); i++ {
		switch rna[i] {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}
	return string(rna)
}
