def solution(n):
    res = 0
    for i in range(n-1):
        if ((i+1) % 3 == 0 or (i+1) % 5 == 0):
            res += i+1
    return res

print(solution(2))
print(solution(10))
print(solution(100))