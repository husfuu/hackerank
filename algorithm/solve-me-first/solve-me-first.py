# def solveMeFirst(a,b):
# 	# Hint: Type return a+b below
#     return a + b


# num1 = int(input())
# num2 = int(input())
# res = solveMeFirst(num1,num2)
# print(res)


def fib(n):
    if n > 1:
        res = fib(n-1) + fib(n-2)
        return res
    elif n == 0:
        return 0
    elif n == 1:
        return 1


res1 = fib(2)
res2 = fib(3)
res3 = fib(4)

print(res1)
print(res2)
print(res3)