package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var DedupLookup map[string]bool

func CombineLists(target, items []string) (int, []string) {
	result := make([]string, len(target))
	copy(result, target)

	for _, item := range result {
		DedupLookup[item] = true
	}

	var totalAdded int
	for _, line := range items {
		if _, ok := DedupLookup[line]; !ok || Force {
			result = append(result, line)
			DedupLookup[line] = true
			totalAdded += 1
		}
	}

	if Sort {
		slices.Sort(result)
	}

	return totalAdded, result
}

func ProcessList(items []string) []string {
	result := make([]string, 0)
	for _, line := range items {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		result = append(result, line)
	}

	if Sort {
		slices.Sort(result)
	}

	return result
}

func ReadList(filename string) ([]string, error) {
	inputFile, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		var e error
		if os.IsNotExist(err) {
			e = fmt.Errorf("file '%s' does not exist", filename)
		} else {
			e = fmt.Errorf("error opening file: %s", err)
		}
		return nil, e
	}
	defer inputFile.Close()

	items := []string{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		items = append(items, scanner.Text())
	}

	return items, nil
}

func ReadAndProcessList(filename string) ([]string, error) {
	list, err := ReadList(filename)
	if err != nil {
		return nil, err
	}

	return ProcessList(list), nil
}

func IsSorted(items []string) bool {
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			return false
		}
	}

	return true
}

func WriteList(filename string, items []string) {
	targetFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Failed to write file")
		return
	}

	if err := targetFile.Truncate(0); err != nil {
		fmt.Printf("error truncating file: %s\n", err)
		return
	}

	for _, item := range items {
		if _, err := targetFile.Write(append([]byte(item), "\n"[0])); err != nil {
			fmt.Println("failed to write file", err)
		}
	}
}

// DEBUG
func debugPrintList(name string, items []string) {
	fmt.Printf("===== %s =====\n", name)
	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println("===== END =====")
}
