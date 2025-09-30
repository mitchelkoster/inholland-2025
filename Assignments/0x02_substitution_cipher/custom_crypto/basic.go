package custom_crypto

func EncodeCipher(input string) string {
	var result []rune // Strings are immutable bytes

	for _, r := range input {
		switch r {
		case 'a':
			result = append(result, 'n')
		case 'b':
			result = append(result, 'o')
		case 'c':
			result = append(result, 'p')
		case 'd':
			result = append(result, 'q')
		case 'e':
			result = append(result, 'r')
		case 'f':
			result = append(result, 's')
		case 'g':
			result = append(result, 'e')
		case 'h':
			result = append(result, 'u')
		case 'i':
			result = append(result, 'v')
		case 'j':
			result = append(result, 'w')
		case 'k':
			result = append(result, 'j')
		case 'l':
			result = append(result, 'y')
		case 'm':
			result = append(result, 'f')
		case 'n':
			result = append(result, 'a')
		case 'o':
			result = append(result, 'b')
		case 'p':
			result = append(result, 'l')
		case 'q':
			result = append(result, 'd')
		case 'r':
			result = append(result, 't')
		case 's':
			result = append(result, 'z')
		case 't':
			result = append(result, 'g')
		case 'u':
			result = append(result, 'h')
		case 'v':
			result = append(result, 'i')
		case 'w':
			result = append(result, 'x')
		case 'x':
			result = append(result, 'k')
		case 'y':
			result = append(result, 'c')
		case 'z':
			result = append(result, 'm')
		default:
			result = append(result, r)
		}
	}
	return string(result)
}

func DecodeCipher(input string) string {
	var result []rune // Strings are immutable bytes

	for _, r := range input {
		switch r {
		case 'n':
			result = append(result, 'a')
		case 'o':
			result = append(result, 'b')
		case 'p':
			result = append(result, 'c')
		case 'q':
			result = append(result, 'd')
		case 'r':
			result = append(result, 'e')
		case 's':
			result = append(result, 'f')
		case 'e':
			result = append(result, 'g')
		case 'u':
			result = append(result, 'h')
		case 'v':
			result = append(result, 'i')
		case 'w':
			result = append(result, 'j')
		case 'j':
			result = append(result, 'k')
		case 'y':
			result = append(result, 'l')
		case 'f':
			result = append(result, 'm')
		case 'a':
			result = append(result, 'n')
		case 'b':
			result = append(result, 'o')
		case 'l':
			result = append(result, 'p')
		case 'd':
			result = append(result, 'q')
		case 't':
			result = append(result, 'r')
		case 'z':
			result = append(result, 's')
		case 'g':
			result = append(result, 't')
		case 'h':
			result = append(result, 'u')
		case 'i':
			result = append(result, 'v')
		case 'x':
			result = append(result, 'w')
		case 'k':
			result = append(result, 'x')
		case 'c':
			result = append(result, 'y')
		case 'm':
			result = append(result, 'z')
		default:
			result = append(result, r)
		}
	}
	return string(result)
}
