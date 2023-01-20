A = [12, 5, 664, 63, 5, 73, 93, 127, 432, 64, 34]
base = 10

for i in range(len(str(max(A)))):
    B = [[] for k in range(base)]
    for x in A:
        figure = x // 10**i % 10
        print(f"Figure from {x} is {figure}")
        B[figure].append(x)
    print(B)
    A = []
    for k in range(base):
        A = A + B[k]
    print(A)
print(A)
