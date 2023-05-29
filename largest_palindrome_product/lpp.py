lst = [1, 2, 3, 4,]
main_list = []

def sorting(lst):
    global main_list
    if len(lst) <= 1:
        return 
    mid_point = len(lst) // 2
    l1 = lst[0:mid_point]
    l2 = lst[mid_point : len(lst)]
    for val1 in l1:
        for val2 in l2:
            main_list.append([val1, val2])
    sorting(l1)
    sorting(l2)
    

print(sorting(lst))
print(main_list)