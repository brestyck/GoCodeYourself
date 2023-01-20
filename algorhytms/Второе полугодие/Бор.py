# Это класс "вершины", или как я его назвал - лист.
class Leaf:
    def __init__(self, character): # Функция инит нужна для создания объекта вершины методом Leaf()
        self.character = character # У вершины есть исходящий из нее "путь" - буква
        self.isLast = False # Здесь указано, является ли вершина концом какого-то слова
        self.number = 0 # Порядковый номер вершины в дереве
        self.leafs = {} # ! Ссылка на другую вершину по следующей букве слова {буква: Leaf}

# Это класс бора (по-английски - Trie)
class Trie(object):
    def __init__(self): # для создания объекта Trie, нужно просто создать одну вершину
        self.initial = Leaf("") # Свойство initial или правильнее root - начало всего дерева, верхняя его точка

    def add(self, word): # Функция добавления в бор слов
        currentLeaf = self.initial # Начинаем добавлять с вершины
        for character in word: # Разбиваем слово на буквы
            if character in currentLeaf.leafs: # Если уже есть ветка с такой буквой...
                currentLeaf = currentLeaf.leafs[character] # Переходим на вершину по этой букве при помощи словаря leafs
            else: # Если нет, то ее надо создать
                new_currentLeaf = Leaf(character) # Вот - создаем
                currentLeaf.leafs[character] = new_currentLeaf # В текущую вершину мы добавляем ссылку на следующую, только
                                                               # что созданную букву
                currentLeaf = new_currentLeaf # И переходим на нее
        currentLeaf.isLast = True # Мы дошли до конца слова, а значит указываем, что последняя буква это конец слова
        currentLeaf.number += 1 # И добавляем +1 к счетчику слов

    # Функция вывода !ОДНОГО СЛОВА! из дерева
    def printout(self):
        lastReached = self.initial.isLast # Аттрибут isLast означает, является ли буква концом слова
        leaf = self.initial # Самая первая вершина
        while not lastReached: # Идем до тех пор, пока не встретим конец слова
            keys = list(leaf.leafs.keys()) # Получаем список ключей словаря
            print(keys[0]) # Из него нас интересует первая буква
            leaf = leaf.leafs[keys[0]] # Мы переходим по первой букве
            lastReached = leaf.isLast # И проверяем, не является ли буква концом слова


if __name__ == "__main__":
    t = Trie()
    t.add("cat")
    t.printout()
