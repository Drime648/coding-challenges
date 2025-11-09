package count


import (
	"io"
	"unicode"
)

func CountBytes(r io.Reader) (int, error) {
	count := 0

	data := make([]byte, 1024)
	for {
		n, err := r.Read(data)
		if (n == 0){
			break
		}
		if err != nil {
			return 0, err
		}
		count += n

	}

	return count, nil
}

func CountLines(r io.Reader) (int, error) {
	count := 0

	data := make([]byte, 1024)

	for {
		n, err := r.Read(data)
		if (n == 0) {
			break
		}
		if err != nil {
			return 0, err
		}

		for i := 0; i < n; i++ {
			if data[i] == '\n' {
				count += 1
			}
		}
	}
	return count, nil
}

func CountWords (r io.Reader) (int, error) {
	count := 0

	data := make([]byte, 1024)

	var prev byte = '\n'
	for {
		n, err := r.Read(data)
		if (n == 0) {
			break
		}
		if err != nil {
			return 0, err
		}
		
		for i := 0; i < n; i++ {
			if unicode.IsSpace(rune(data[i])) && !unicode.IsSpace(rune(prev)){
				count += 1
			}
			prev = data[i]
		}
	}
	return count, nil
}
