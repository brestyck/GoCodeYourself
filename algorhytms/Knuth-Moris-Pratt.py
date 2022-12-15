def prefix(text):
    result = [0]
    for i in range(1, len(text)):
        counter = 0
        prefixBorder = i
        while prefixBorder < len(text):
            if text[counter] == text[prefixBorder]:
                prefixBorder += 1
                counter += 1
            else:
                break

        result.append(counter)
    return result


def kmp(text, pattern):
    text = pattern + "\u0002" + text
    return prefix(text), text


print(kmp("Whatsoever", "so"))
