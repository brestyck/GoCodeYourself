# Два пузырька бегают туда-сюда и сортируют
# внутри границ end_index start_index
# Когда границы схлопываются, то все

def cocktailSort(a):
    start_index = 0
    end_index = len(a) - 1
    while True:
        for i in range(start_index, end_index):
            if a[i] > a[i + 1]:
                a[i], a[i + 1] = a[i + 1], a[i]
        end_index -= 1
        
        if end_index == start_index:
            break
        
        for i in range(end_index-1, start_index-1, -1):
            if a[i] > a[i + 1]:
                a[i], a[i + 1] = a[i + 1], a[i]
        start_index += 1


a = [14, 51, 0, 2, 3, 1]
cocktailSort(a)
print(f"Sorted array is: {a}")
