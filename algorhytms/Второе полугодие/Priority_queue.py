class Node:
    val = None
    linkToNext = None
    def __init__(self, val, linkToNext = None, linkToPrev = None) -> None:
        self.val = val
        self.linkToNext = linkToNext
        self.linkToPrev = linkToPrev

class AlgorhytmQueue:
    head = None
    def __init__(self) -> None:
        pass
    # Функция добавления
    def append(self, val) -> None:
        # Если не существует начало, создаем его и вылетаем
        if not self.head:
            self.head = Node(val)
            return val
        current = self.head
        while current.linkToNext:
            current = current.linkToNext
        # У конечного вставляем ссылку на предыдущий
        current.linkToNext = Node(val, linkToPrev=current)
    # Функция вывода в консоль
    def printout(self):
        current = self.head
        print("[", end="")
        while current.linkToNext:
            print(current.val, end=" ")
            current = current.linkToNext
        print(current.val, end="")
        print("]")
    # Получение по индексу
    def getById(self, id: int) -> None:
        i = 0
        current = self.head
        while i < id:
            if current.linkToNext:
                current = current.linkToNext
            else:
                raise IndexError
            i += 1
        return current.val
    # Удаление по индексу
    def delById(self, id: int) -> None:
        current = self.head
        i = 0
        if id == 0:
            self.head = self.head.linkToNext
            return
        while current.linkToNext:
            if i != id:
                current = current.linkToNext
            else:
                if current.linkToNext:
                    current.linkToNext.linkToPrev = current.linkToPrev
                if current.linkToPrev:
                    current.linkToPrev.linkToNext = current.linkToNext
                del current
                break
            i += 1
    # Перезапись
    def reWrite(self, id: int, val: any) -> None:
        current = self.head
        i = 0
        while current.linkToNext:
            if i == id:
                current.val = val
            current = current.linkToNext
            i += 1
    def add(self, id: int, val: any) -> None:
        current = self.head
        i = 0
        if id == 0:
            self.head = Node(val, linkToNext=self.head)
            return
        while current.linkToNext:
            if i == id:
                current.linkToPrev.linkToNext = Node(val, linkToNext=current, linkToPrev=current.linkToPrev)
            current = current.linkToNext
                
            i += 1

if __name__ == "__main__":
    test_queue = AlgorhytmQueue()
    test_queue.append(1)
    test_queue.append(213213)
    test_queue.append(45)
    test_queue.append(23)
    test_queue.printout()
    test_queue.add(0, 666)
    test_queue.printout()
    test_queue.delById(3)
    test_queue.printout()
    test_queue.reWrite(2, 0)
    test_queue.printout()
    print(test_queue.getById(0))
    