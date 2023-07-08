x = 1 # 0, 1 factorial
y = 1 # 0, 1 factorial
z = 2 # 0, 1, 2 factorial
n = 3

""" looping standar
"""
# res = []
# for i in range(0, x+1):
#     for j in range(0, y+1):
#         for z in range(0, z+1):
#             if (i + j + z != n):
#                 res.append([i, j, z])

# print(res)

"""list comprehensions
"""
mylist = []
res2 = [mylist+[i,j,k] for i in range(0,x+1) for j in range(0,y+1) for k in range(0,z+1) if  i + j +k != n]

print(res2)

# if __name__ == '__main__':
#     x = int(input())
#     y = int(input())
#     z = int(input())
#     n = int(input())
