# Memory Database Sorting  

This project implements a **singly linked list–based memory database** that can load student records from a CSV file, store them in memory as linked list nodes, and then perform sorting operations using two classical algorithms: **Bubble Sort** and **Insertion Sort**.  

The database is implemented in both **Python** and **Go**, demonstrating how the same logic can be applied across programming languages while managing nodes directly, without converting the data into arrays.  

---

## 🚀 Features  
- **Recursive Data Loading**  
  Student records are read from a CSV file and inserted into a singly linked list one node at a time using recursion.  

- **Sorting Algorithms**  
  - **Bubble Sort**: Sorts the linked list by repeatedly swapping adjacent nodes.  
  - **Insertion Sort**: Builds a sorted sublist by inserting nodes into their correct position.  
  Sorting can be done on any column such as `age` or `absences`.  

- **Recursive Export to CSV**  
  After sorting, results are written back to new CSV files using a recursive function that writes one node per row.  
  - Bubble Sort results → `bubble_student-data.csv`  
  - Insertion Sort results → `insertion_student-data.csv`  

- **Cross-Language Implementation**  
  - **Python**: Easy-to-read implementation with classes, recursion, and CSV handling.  
  - **Go**: High-performance implementation with maps, structs, and slices to manage nodes.  

---

## 🗂 Project Structure  

memory-database-sorting/
│── python/
│ ├── memory_database.py # Python implementation
│ ├── student-data.csv # Input dataset
│ ├── bubble_student-data.csv # Bubble sorted output
│ ├── insertion_student-data.csv# Insertion sorted output
│
│── go/
│ ├── memory_database.go # Go implementation
│ ├── student-data.csv # Input dataset
│ ├── bubble_student-data.csv # Bubble sorted output
│ ├── insertion_student-data.csv# Insertion sorted output
│
│── README.md # Project documentation

---

## ⚙️ How to Run  

### Python  
```bash
cd python
python3 memory_database.py
```
Loads student-data.csv
Performs Bubble Sort (by age) and Insertion Sort (by absences)
Exports results to bubble_student-data.csv and insertion_student-data.csv

Go bash
```
cd go
go run memory_database.go
```
Loads student-data.csv
Performs Bubble Sort (by age) and Insertion Sort (by absences)
Exports results to bubble_student-data.csv and insertion_student-data.csv

📊 Performance Comparison
System performance was monitored using Linux commands like top and ps.
Both CPU and disk usage were observed while executing Bubble Sort and Insertion Sort.
Bubble Sort requires more CPU cycles due to repeated comparisons and swaps.
Insertion Sort uses fewer comparisons and tends to perform better on partially sorted data.

📌 Example Output
Terminal (Python):

Original data from CSV file
Total nodes: 396
After Bubble Sort by age
Exported bubble sorted csv file named : bubble_student-data.csv
After Insertion Sort by absences
Exported insertion sorted csv file named : insertion_student-data.csv

CSV Output (insertion_student-data.csv):
school,sex,age,address,famsize,Pstatus,Medu,Fedu,Mjob,Fjob,reason,guardian,...
GP,F,15,U,GT3,T,4,4,teacher,services,home,mother,...,absences:0
GP,F,16,U,LE3,T,2,2,other,other,reputation,mother,...,absences:2

...
🙏 Acknowledgment
This project was developed as part of coursework under the guidance of my Professor, whose lectures and assignments inspired the exploration of algorithms within memory-based data structures. Special thanks to the professor for their continuous support, valuable insights, and encouragement throughout the project.

📎 License
This project is shared for educational purposes. You are free to fork and use it with proper attribution.

---
