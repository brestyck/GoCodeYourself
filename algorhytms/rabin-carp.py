temporary_cache = 0
primeConstant = 5179
alphabetPower = 256


def calculate_hash(string):
    cache = 0
    n = len(string)
    for i in range(n, 0, -1):
        power = i-1
        cache += (ord(string[n-i])*(alphabetPower**power)) % primeConstant
    return cache % primeConstant


def find_occur(text, pattern):
    patternLength = len(pattern)
    patternCache = calculate_hash(pattern)
    counter = 0
    print(f"Searching for {patternCache} ({pattern})")

    initCache = calculate_hash(text[0:patternLength])
    if initCache == patternCache:
        print("Cache match found")
        if text[0:patternLength] == pattern:
            print("Full match")
            counter += 1
    print(f"Initial is {initCache} for {text[0:patternLength]}")

    for j in range(patternLength-1, len(text)-1):
        power = patternLength - 1
        initCache = initCache - ord(text[j-(patternLength-1)])*alphabetPower**power  # -a*d^2
        initCache *= alphabetPower  # polynom * d
        initCache += ord(text[j+1])  # +x*d^0
        initCache %= primeConstant
        # print(initCache)
        if initCache < 0:
            initCache += primeConstant
        print(f"Cache of substring {text[j-(patternLength-1)+1:j+2]} is {initCache}")
        if initCache == patternCache:
            print("Cache match found")
            if text[j-(patternLength-1)+1:j+2] == pattern:
                print("Full match")
                counter += 1
    return counter


print(find_occur("Pycharm is hardly b", "Py"))
# print(find_occur("abbaba", "bab"))
