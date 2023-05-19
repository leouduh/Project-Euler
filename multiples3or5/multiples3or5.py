def multiples3or5(num):
    sum = 0
    for i in range(3, num):
        if i % 3 == 0:
            sum += i
        elif (i % 5 == 0) and i >= 5:
            sum += i
    return sum

num = 1000
print(f'Sum of all multiples of 3 and 5 from 0 to {num} is {multiples3or5(num)}')