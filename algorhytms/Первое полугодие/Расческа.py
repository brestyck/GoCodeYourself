from random import randint

list_1 = [randint(1,100) for a in range(10)]
n = len(list_1)

step = n

while step > 1 or q:
    if step > 1:
        step -= 3
    q, i = False, 0
    while i + step < n:
        if list_1[i] > list_1[i + step]:
            list_1[i], list_1[i + step] = list_1[i + step], list_1[i]
            print(list_1)
            q = True
        i += step

print(list_1)
