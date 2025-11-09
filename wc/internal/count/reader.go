package count


import (
	"io"
	"unicode"
)

type Stats struct {
	NumBytes int
	NumWords int
	NumLines int
	NumChars int
}

func CountData(r io.Reader) (Stats, error) {

	stats := Stats {
		NumBytes: 0,
		NumWords: 0,
		NumLines: 0,
		NumChars: 0,
	}

	data := make([]byte, 1024)

	var prev byte = '\n'
	for {
		n, err := r.Read(data)
		if (n == 0) {
			break
		}
		if err != nil {
			return Stats{}, err
		}
		stats.NumBytes += n

		stringData := string(data[:n])
		for range stringData {
			stats.NumChars++  // using r implicitly for counting
		}
		
		for i := 0; i < n; i++ {
			if unicode.IsSpace(rune(data[i])) && !unicode.IsSpace(rune(prev)){
				stats.NumWords += 1
			}
			if data[i] == '\n' {
				stats.NumLines += 1
			}
			prev = data[i]
		}
	}
	return stats, nil
}






// func CountBytes(r io.Reader) (int, error) {
// 	count := 0
//
// 	data := make([]byte, 1024)
// 	for {
// 		n, err := r.Read(data)
// 		if (n == 0){
// 			break
// 		}
// 		if err != nil {
// 			return 0, err
// 		}
// 		count += n
//
// 	}
//
// 	return count, nil
// }
//
// func CountLines(r io.Reader) (int, error) {
// 	count := 0
//
// 	data := make([]byte, 1024)
//
// 	for {
// 		n, err := r.Read(data)
// 		if (n == 0) {
// 			break
// 		}
// 		if err != nil {
// 			return 0, err
// 		}
//
// 		for i := 0; i < n; i++ {
// 			if data[i] == '\n' {
// 				count += 1
// 			}
// 		}
// 	}
// 	return count, nil
// }
//
// func CountWords (r io.Reader) (int, error) {
// 	count := 0
//
// 	data := make([]byte, 1024)
//
// 	var prev byte = '\n'
// 	for {
// 		n, err := r.Read(data)
// 		if (n == 0) {
// 			break
// 		}
// 		if err != nil {
// 			return 0, err
// 		}
//
// 		for i := 0; i < n; i++ {
// 			if unicode.IsSpace(rune(data[i])) && !unicode.IsSpace(rune(prev)){
// 				count += 1
// 			}
// 			prev = data[i]
// 		}
// 	}
// 	return count, nil
// }

