# Сделано Денисом Степановым и для меня здесь нихера не понятно

def build_pat(pattern: str):
    table = {}
    ind = 1
    for i in range(len(pattern) - 2, -1, -1):
        if pattern[i] not in table:
            table[pattern[i]] = ind
        else:
            table[pattern[i]] = table[pattern[i]]
        ind += 1
    table[pattern[-1]] = len(pattern)
    return table

def find(stroka: str, pattern: str):
    start = None
    end = None
    table = build_pat(pattern)
    mark = len(pattern) - 1
    while mark < len(stroka):
        if pattern[-1] == stroka[mark]:
            flag = True
            ch = mark
            for i in range(len(pattern) - 1, -1, -1):
                if pattern[i] != stroka[ch]:
                    mark += table[pattern[i]]
                    flag = False
                    break
                ch -= 1
            if flag is True:
                start = mark - len(pattern) + 1
                end = mark
                break
        else:
            if stroka[mark] in table:
                mark += table[stroka[mark]]
            else:
                mark += len(pattern)
    return start, end



stroka = input()
pattern = input()
ind = find(stroka, pattern)
print(ind)
print(stroka[ind[0]:ind[1] + 1])