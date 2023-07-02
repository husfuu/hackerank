class SinglyLinkedListNode:
    def __init__(self, node_data):
        self.data = node_data
        self.next = None

class SinglyLinkedList:
    def __init__(self):
        self.head = None

def print_singly_linked_list(node, sep, fptr):
    while node:
        fptr.write(str(node.data))

        node = node.next

        if node:
            fptr

def printLinkedList(head):
    # head: SinglyLinkedListNode | self.head 
    i = 0
    while True:
        i += 1
        if head == None:
            break
        print(head.data)
        head = head.next


# def insertNodeAtTail(head, data):
#     node = SinglyLinkedListNode(data)
#     if head != None:
#         # print('menuhin')
#         head.next = node
#     return node
#     # if head == None:
#     #     node.data = data
#     # else:
#     #     head.next = node

def insertNodeAtTail(head, data):
    node = SinglyLinkedListNode(data)
    if head == None:
        # skenario jika Node(head) bernilai null, maka dibikin node baru
        return node
    # skenario jika Node(head) tidak null
    # dicari dulu Node yang paling ujung dari Node yang diberikan, jika udah dapet baru dimasukan
    # nilai next-nya -> node_paling_ujung.next = node_baru
    curr = head
    while True:
        if curr.next == None:
            break
        curr = curr.next

    curr.next = node
    return head


llist = SinglyLinkedList()

llist_head = insertNodeAtTail(llist.head, 142)
llist.head = llist_head

llist_head = insertNodeAtTail(llist.head, 302)
llist.head = llist_head



print('masuk')
printLinkedList(llist.head)
print('masuk')





# TESTING AJAH
llist = SinglyLinkedList()
node1 = SinglyLinkedListNode(141)
node2 = SinglyLinkedListNode(302)

node1.next = node2
llist.head = node1

printLinkedList(llist.head)



if __name__ == '__main__':
    llist_count = int(input())

    llist = SinglyLinkedList() # data linkedlist = null atau None

    for i in range(llist_count):
        llist_item = int(input())
        llist_head = insertNodeAtTail(llist.head, llist_item)
        llist.head = llist_head

    printLinkedList(llist.head)
