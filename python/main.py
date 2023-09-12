num1 = [1, 1, 2]
num2 = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]


# def solution(num):
#     res = []
#     for i in range(len(num1)):
#         for j in range(len(num1)-i):
#             print(f"{j-i} | {len(num1)-i}")
#             if num1[i] == num1[j-i]:
#                 continue
#         res.append(num1[i])

#     return res


# for i in range(len(num1)):
#     for j in range(len(num1)-i-1):
#         y = j + i + 1
#         print(f"y : {num1[y]} | i : {num1[i]}")
#     print("-----")


def solution11(num):
    res = []
    for i in range(len(num)):
        same = False
        for j in range(len(num)-i-1):
            j = j + i + 1
            if num[i] == num[j]:
                same = True
        if same != True:
            count += 1
            res.append(num[i])
    return res



def solution(nums):
    count = 0
    for i in range(len(nums)):
        same = False
        for j in range(len(nums)-i-1):
            j = j + i + 1
            if nums[i] == nums[j]:
                same = True
        if same != True:
            count += 1
    return count


# def removeDuplicates(nums):
#     i,j=0,1
#     while i<=j and j<len(nums):
#         if nums[i]==nums[j]:
#             del nums[j]
#         else:
#             i+=1
#             j+=1
#     return len(nums)


# def removeDuplicates(nums):
#     replace = 1
#     for i in range(1, len(nums)):
#         if nums[i-1] != nums[i]:
#             nums[replace] = nums[i]
#             replace += 1
#     return replace



# # res = solution(num1)
# # res = solution(num2)
# res = removeDuplicates(num1)

# print(res)


def fizzBuzz(n):
    res = []
    for i in range(1, n+1):
        if (i % 3 == 0) and (i % 5 == 0):
            res.append("FizzBuzz")
        elif i % 3 == 0:
            res.append("Fizz")
        elif i % 5 == 0:
            res.append("Buzz")
        else:
            print('bangke')
            res.append(f"{i}")
    return res


res = fizzBuzz(3)
print(res)