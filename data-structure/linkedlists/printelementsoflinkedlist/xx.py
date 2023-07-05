def sum_mul(n, m):
    pass

n = 2
m = 9
sum_i = 0
for i in range(m):
    temp = (i+1) * n
    if temp >= m:
        break
    sum_i += temp
print(sum_i)




def sum_mul(n, m):
    if m <= 0 or n <= 0:
        return 'INVALID'
    
    sum_i = 0
    for i in range(m):
        temp = (i+1)*n
        if temp >= m:
            break
        sum_i += temp
    return sum_i 

res = sum_mul(2, 9)

print(res)