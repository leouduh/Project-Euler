def fibonacci(end_point: int)-> tuple:
    arr = [1,]
    x = 1
    y = 2
    sum = 0
    while y < end_point:
        arr.append(y)
        if y%2 == 0:
            sum += y
        x, y = y, x + y
    return arr, sum

series, sum = fibonacci(4_000_000)
print(f"The series is {series}")
print(f"The sum of all even fibonacci terms in the series above is {sum}")
