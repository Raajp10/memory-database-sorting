package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// create node 
type Node struct {
	// store one row and pointer to next node
	list map[string]string
	next *Node
}

// make single link list database
type SingleLinkedListDatabase struct {
	head *Node
	size int
}

// add a new node at the end for that we take input as current node and data
func (db *SingleLinkedListDatabase) add_data(data map[string]string) {
	newNode := &Node{list: data, next: nil}
	if db.head == nil {
		db.head = newNode
	} else {
		n := db.head
		for n.next != nil {
			n = n.next
		}
		n.next = newNode
	}
	db.size++
}

// load data row by row into linklist and default index value is 0
func (db *SingleLinkedListDatabase) load_data(rows []map[string]string, i int) {
	// stop if index is greater or equal to number of rows then return
	if i >= len(rows) {
		return
	}
	// add row to linklist and increase index to one for next
	db.add_data(rows[i])
	db.load_data(rows, i+1)
}

// print database loaded in link list
func (db *SingleLinkedListDatabase) print_data() {
	// start from head node and find until next is nil
	n := db.head
	for n != nil {
		fmt.Println(n.list)
		n = n.next
	}
}

// cast funnction helps values for comparisons so it never mix str & numbers
func (db *SingleLinkedListDatabase) cast(value string) (int, float64, string) {
	s := strings.TrimSpace(value)
	//  0 for float 
	if s == "" {
		return 1, 0, ""
	}
	// 1 for string 
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return 0, f, ""
	}
	return 1, 0, strings.ToLower(s)
}

// bubble sort
func (db *SingleLinkedListDatabase) bubble_sort(key string) {
	// check if head is empty 0 or 1 node then return
	if db.head == nil || db.head.next == nil {
		return
	}

	swapped := true
	for swapped {
		swapped = false
		current := db.head
		for current != nil && current.next != nil {
			// convert values to a comparable form
			tag1, num1, str1 := db.cast(current.list[key])
			tag2, num2, str2 := db.cast(current.next.list[key])

			// if value1 is greater then swap
			greater := false
			if tag1 > tag2 {
				greater = true
			} else if tag1 == tag2 {
				if tag1 == 0 { // numeric
					if num1 > num2 {
						greater = true
					}
				} else { // string
					if str1 > str2 {
						greater = true
					}
				}
			}

			if greater {
				current.list, current.next.list = current.next.list, current.list
				swapped = true
			}
			current = current.next
		}
	}
}

// insertion sort
func (db *SingleLinkedListDatabase) insertion_sort(key string) {
	sorted_head := (*Node)(nil)
	current := db.head
	// take nodes one by one from the list and insert into the sorted sublist
	for current != nil {
		// save the remainder before we connect it
		next_node := current.next
		sorted_head = db.insert_sorted(sorted_head, current, key)
		current = next_node
	}
	db.head = sorted_head
}

func (db *SingleLinkedListDatabase) insert_sorted(head *Node, node *Node, key string) *Node {
	// detach node before inserting
	node.next = nil

	tagN, numN, strN := db.cast(node.list[key])

	// if the sorted list is empty OR node belongs at the front
	if head == nil {
		node.next = head
		return node
	}
	tagH, numH, strH := db.cast(head.list[key])
	atFront := false
	if tagN < tagH {
		atFront = true
	} else if tagN == tagH {
		if tagN == 0 { // numeric
			if numN <= numH {
				atFront = true
			}
		} else { // string
			if strN <= strH {
				atFront = true
			}
		}
	}
	if atFront {
		node.next = head
		return node
	}

	// walk until we find the first element that is >= node
	prev, current := head, head.next
	for current != nil {
		tagC, numC, strC := db.cast(current.list[key])

		goOn := false
		if tagN > tagC {
			goOn = true
		} else if tagN == tagC {
			if tagN == 0 {
				if numN > numC {
					goOn = true
				}
			} else {
				if strN > strC {
					goOn = true
				}
			}
		}
		if !goOn {
			break
		}
		prev, current = current, current.next
	}

	// splice node between previous and current
	prev.next, node.next = node, current
	return head
}

// export csv file from loaded linklist
func (db *SingleLinkedListDatabase) export_to_csv_file(path string, data []string) error {
	// drop the 'key' column that is for reference to access node
	exportData := make([]string, 0, len(data))
	for _, h := range data {
		if h != "key" {
			exportData = append(exportData, h)
		}
	}

	// write header first
	if err := ReadAndWriteCSV_write_csv(path, exportData); err != nil {
		return err
	}

	// then write rows by recursively until end
	var write_node func(*Node) error
	write_node = func(n *Node) error {
		if n == nil {
			return nil
		}
		// write file in append mode line by line
		f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		w := csv.NewWriter(f)
		row := make([]string, len(exportData))
		for i, h := range exportData {
			row[i] = n.list[h]
		}
		if err := w.Write(row); err != nil {
			_ = f.Close()
			return err
		}
		w.Flush()
		if err := w.Error(); err != nil {
			_ = f.Close()
			return err
		}
		_ = f.Close()
		// move to the next node in the linked list
		return write_node(n.next)
	}

	// start from the actual head (we did NOT inject a header node)
	start := db.head
	return write_node(start)
}

// read csv opens CSV and write
func ReadAndWriteCSV_read_csv(path string) ([]map[string]string, []string, error) {
	rows := []map[string]string{}
	var headers []string

	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	all, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	if len(all) == 0 {
		return rows, headers, nil
	}
	headers = all[0]

	// add an indexing key for identifying row (as string)
	for i, rec := range all[1:] {
		row := map[string]string{"key": fmt.Sprintf("%d", i+1)}
		for j, h := range headers {
			if j < len(rec) {
				row[h] = rec[j]
			} else {
				row[h] = ""
			}
		}
		rows = append(rows, row)
	}

	// ensure "key" is included in headers at front (like Python)
	hasKey := false
	for _, h := range headers {
		if h == "key" {
			hasKey = true
			break
		}
	}
	if !hasKey {
		headers = append([]string{"key"}, headers...)
	}

	return rows, headers, nil
}

// write_csv writes a single header row to path (overwrites file)
func ReadAndWriteCSV_write_csv(path string, headers []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err := w.Write(headers); err != nil {
		return err
	}
	w.Flush()
	return w.Error()
}

// main 
func main() {
	// read csv file rows is data and header is all column name
	rows, header, err := ReadAndWriteCSV_read_csv("./student-data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// create an empty linked list database
	db := &SingleLinkedListDatabase{}
	// recursive load data
	db.load_data(rows, 0)

	// print data
	fmt.Println("Original data from CSV file")
	fmt.Printf("\nTotal nodes: %d\n", db.size+1)
	db.print_data()

	// bubble sort by "age" column
	db.bubble_sort("age")
	fmt.Println("After bubble Sort by age")
	db.print_data()
	fmt.Printf("\nTotal nodes: %d\n", db.size+1)

	// export loaded data into csv file
	if err := db.export_to_csv_file("bubble_student-data.csv", header); err != nil {
		fmt.Println("Export error:", err)
		return
	}
	fmt.Println("Exported bubble sorted csv file named : bubble_student-data.csv")

	// create a new linked list database for insertion sort
	db2 := &SingleLinkedListDatabase{}
	db2.load_data(rows, 0)
	// insertion sort by "absences" column
	db2.insertion_sort("absences")
	fmt.Println("After Insertion Sort by absences")
	db2.print_data()
	fmt.Printf("\nTotal nodes: %d\n", db2.size+1)

	// export loaded data into csv file
	if err := db2.export_to_csv_file("insertion_student-data.csv", header); err != nil {
		fmt.Println("Export error:", err)
		return
	}
	fmt.Println("Exported insertion sorted csv file named : insertion_student-data.csv")
}
