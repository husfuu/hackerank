n1 = -14000050000


# def no_boring_zeros(n):
#     # your code
#     new_n = ""
#     i = 0
#     for digit in str(n):
#         if digit == '0' and i != 0:
#             return int(new_n)
#         new_n += digit
#         i += 1
#     return int(new_n)


# nums_str = str(n1)
# bakward_str = range(len(nums_str)+1)[::-1]
# print("bakward_str: ", bakward_str)
# for i in bakward_str:
#     print("xx: ", i)



def no_boring_zeros(n):
    nums_str = str(n)
    N = len(nums_str)
    new_n = ""
    for i in range(N, -1, -1):
        if N == 1 and nums_str[i-1] == '0':
            print("i: ", N)
            return int(0)
        if nums_str[i-1] != '0':
            new_n = nums_str[:i] + new_n
            break
    
    return int(new_n)


n1 = 0
res = no_boring_zeros(n1)
# res = no_boring_zeros(0)
print("hasil: ", res)

print("hasil: ", type(res))
