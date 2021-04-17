package main
    
import (
	"github.com/tsak/concurrent-csv-writer"
	"strconv"
)

func main() {
	// Create `sample.csv` in current directory
	csv, err := ccsv.NewCsvWriter("sample.csv")
	if err != nil {
		panic("Could not open `sample.csv` for writing")
	}

	// Flush pending writes and close file upon exit of main()
	defer csv.Close()

	count := 10000

	done := make(chan bool)

	for i := count; i > 0; i-- {
		go func(i int) {
			csv.Write([]string{strconv.Itoa(i), "bottles", "of", "beer"})
			done <- true
			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	for i := 0; i < count; i++ {
		<-done
	}
}



