def selection_sort(a):
    for i in range(0, len(a)-1):
        for j in range(i+1, len(a)):
            if a[j] < a[i]:
                a[j], a[i] = a[i], a[j]
    return a

array = [51, 9, 0, 2, -1, 43]
print(selection_sort(array))
