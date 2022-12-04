# Грубо говоря, двигаем отсортированную границу влево
# На несорт. участке справа ищем мин. число
# Мин число (min_index) меняем местами с текущей границей (i)

def insertionSort(array):
	for i in range(1, len(array)):
		key = array[i]
		j = i - 1
		while j >= 0 and key < array[j]:
			array[j + 1] = array[j]
			j -= 1
		array[j + 1] = key
		print(*array)


a = [12, 11, 13, 5, 6]
insertionSort(a)
print(*a)

# СДЕЛАТЬ РАСЧЕСКОЙ И ВЫБОРОМ
