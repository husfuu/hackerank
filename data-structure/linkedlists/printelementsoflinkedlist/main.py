class SinglyLinkedListNode:
    def __init__(self, node_data):
        self.data = node_data
        self.next = None

class SinglyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
    def insert_node(self, node_data):
        node = SinglyLinkedListNode(node_data)
        if not self.head:
            # jika self.head == none
            self.head = node
        else: # jika self.head != none, maka tambahkan "next"-nya
            self.tail.next = node
        
        self.tail = node

def printLinkedList(head):
    # head: SinglyLinkedListNode | self.head 
    i = 0
    while True:
        i += 1
        if head == None:
            break
        print(head.data)
        head = head.next


if __name__ == '__main__':
    llist_count = int(input())    
    llist = SinglyLinkedList()

    for _ in range(llist_count):
        llist_item = int(input())
        llist.insert_node(llist_item)

    printLinkedList(llist.head)
