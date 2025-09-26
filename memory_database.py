import csv
# create nodes for single link list 
class Node:
    def __init__(self,list):
        # store one row from and next node is point to none 
        self.list = list
        self.next= None

#create database using single linklist
class SingleLinkedListDatabase:
    def __init__(self):
        # start with an empty single linklist here head stores row ,size for length
        self.head = None
        self.size = 0

    # add a new node at the end for that we take input as current node and data 
    def add_data(self, data):
        # make new node and store data
        new_node = Node(data)
        # if the list is completely empty then add head
        if self.head is None:
            self.head = new_node
        # add last node's next new node and new's next is none
        else:
            n = self.head
            while n.next is not None:
                n = n.next
            n.next = new_node
        self.size += 1
    
    # load data row by row into linklist and default index value is 0
    def load_data(self,rows,i=0):
        #  stop if index is greater or equal to number of rows then return
        if i >= len(rows):
            return
        # add row to linklist and increase index to one for next 
        self.add_data(rows[i])
        self.load_data(rows, i + 1)
    
    # print database loaded in link list
    def print_data(self):
        # start from head node and find until next is none
        n=self.head
        while n is not None:
            print(n.list)
            n = n.next

    # cast funnction helps values for comparisons so it never mix str & numbers
    def cast(self, value):
        s = "" if value is None else str(value).strip()
        try:
            # 0 for float 
            return (0, float(s))
        except ValueError:
            # 1 for string
            return (1, s.lower())


    # bubble sort 
    def bubble_sort(self, key):
        # check if head is empty (0 or 1 node) then return
        if not self.head or not self.head.next:
            return

        swapped = True
        while swapped:
            swapped = False
            current = self.head
            while current and current.next:
                # convert values to a comparable form 
                value1 = self.cast(current.list.get(key))
                value2 = self.cast(current.next.list.get(key))

                # if value1 is greater then swap 
                if value1 > value2:
                    current.list, current.next.list = current.next.list, current.list
                    swapped = True

                current = current.next
   
    # insertion sort
    def insertion_sort(self, key):
        sorted_head = None
        current = self.head
        # take nodes one by one from the list and insert into the sorted sublist
        while current:
            # save the remainder before we connect it
            next_node = current.next
            sorted_head = self.insert_sorted(sorted_head, current, key)
            current = next_node
        self.head = sorted_head


    def insert_sorted(self, head, node, key):
        # detach node before inserting
        node.next = None

        node_key = self.cast(node.list.get(key))

        # if the sorted list is empty OR node belongs at the front
        if head is None or node_key <= self.cast(head.list.get(key)):
            node.next = head
            return node

        # walk until we find the first element that is >= node
        prev, current = head, head.next
        while current and node_key > self.cast(current.list.get(key)):
            prev, current = current, current.next

        # splice node between previous and current
        prev.next, node.next = node, current
        return head
    
    # export csv file from loaded linklist
    def export_to_csv_file(self, path, data):
        # drop the 'key' column that is for reference to access node
        export_data = [h for h in data if h != "key"]

        # write header first  
        ReadAndWriteCSV.write_csv(path, export_data)

        # then write rows by recursively until end
        def write_node(n):
            if n is None:
                return
            # write file in append mode line by line
            with open(path, "a", newline="") as f:
                w = csv.writer(f)
                w.writerow([n.list.get(h, "") for h in export_data])
            # move to the next node in the linked list
            write_node(n.next)
        # skip the header node to avoid repeation
        # start = self.head.next if self.head else None  
        start = self.head
        write_node(start)
    
    
# read and write for csv file for that make class
class ReadAndWriteCSV:
    def read_csv(path):
        # rows for data and headers for column name
        rows = []
        headers = None
        # open the CSV file and read line as  dictionary and first line for header 
        with open(path, newline="") as f:
            reader = csv.DictReader(f)
            headers = reader.fieldnames
             # add a indexing key for identify row 
            i = 1
            for r in reader:
                r = dict(r)
                r["key"] = str(i)
                rows.append(r)
                i += 1
        # check that key is there or not 
        if "key" not in headers:
            headers = ["key"] + headers 
        # return the list that contain data
        return rows, headers
    
    def write_csv(path, headers):
        # open file in write mode and write line by line  
        with open(path, "w", newline="") as f:
            w = csv.writer(f)
            w.writerow(headers)
    
# main function 
def main():    
     # read csv file rows is data and header is all column name 
    rows, header = ReadAndWriteCSV.read_csv("./student-data.csv")
    # create an empty linked list database
    db = SingleLinkedListDatabase()
    # recursive load data
    db.load_data(rows, 0)
    # print data
    print("Original data from CSV file")
    # database size + header fileds
    print(f"\nTotal nodes: {db.size+1}")
    db.print_data()

    # bubble sort by "age" column
    db.bubble_sort("age")
    print("After bubble Sort by age")
    db.print_data()
    print(f"\nTotal nodes: {db.size+1}")

    # export loaded data into csv file
    db.export_to_csv_file("bubble_student-data.csv", header)
    print("Exported bubble sorted csv file named : bubble_student-data.csv")

    # create an empty linked list database
    db2 = SingleLinkedListDatabase()
    # recursive load data
    db2.load_data(rows, 0)
    # insertion sort by "absences" column
    db2.insertion_sort("absences")
    print("After Insertion Sort by absences")
    db2.print_data()
    print(f"\nTotal nodes: {db.size+1}")

     # export loaded data into csv file 
    db2.export_to_csv_file("insertion_student-data.csv", header)
    print("Exported insertion sorted csv file named : insertion_student-data.csv")

    print(f"\nTotal nodes: {db.size+1}")

if __name__ == "__main__":
    main()
