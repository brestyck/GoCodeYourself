# Создаем словарь, содержащий сдвиги
def createSetAndDict(substr: str) -> dict:
    dct = {}
    uniques = set()
    # Перебираем буквы от предпоследней до первой
    for letter_i in range(len(substr)-2, -1, -1):
        letter = substr[letter_i]
        # Закидываем в сет и словарь
        if not letter in uniques:
            dct[letter] = (len(substr)-1)-letter_i # Расстояние до конца (симв.)
            uniques.add(letter)
        print(f"Iterating {letter_i} -> {letter}, exists in set: {letter in uniques}")
    # В конце добавляем для последней буквы
    dct[substr[-1]] = len(substr)
    uniques.add(substr[-1])
    print(f"Created the following dictionary {dct}")
    print(f"With the following set: {uniques}")
    return dct, uniques

# Функция поиска подстроки
def findSubstringIndexes(string: str, substr: str, shifts: dict, uniques: set) -> int:
    substrLength = len(substr)
    searchContinues = True # have we found?
    pointer = substrLength - 1 # where are we?
    while searchContinues:
        print(f"Индекс: {pointer}")
        if pointer > len(string):
            searchContinues = False
            break
        letter = string[pointer] # прислоняем
        if letter == substr[-1]: # Если совпало с последней, то ...
            subpointer = pointer
            fullMatch = True
            for i in range(len(substr) - 1, -1, -1): # Проверям, а не совпадает ли с остальными
                if substr[i] != string[subpointer]:
                    pointer += shifts[substr[i]] # Сдвигаемся именно по СЛОВАРЮ
                    fullMatch = False
                    break
                subpointer -= 1 # Идем по ней назад
            if fullMatch:
                # Индексы подстроки в строке
                return pointer - len(substr) + 1, pointer + 1
        else: # А если не совпало, идем по большому сдвигу
            if letter in uniques:
                pointer += shifts[letter]
            else:
                pointer += substrLength
    # Если ничего не нашли, возвращаем 0 0
    return 0, 0


if __name__ == "__main__":
    string = "abracadabra"
    substr = "racad"
    shifts, uniques = createSetAndDict(substr)
    print("Поиск...")
    i, j = findSubstringIndexes(string, substr, shifts, uniques)
    print(f"Найдено: {i}, {j}")
    print(f"{string[i:j]}")
    print(f"")
    print(f"")
