import sys
sys.setrecursionlimit(30000)

def neighbours(p:(int,int)) -> list:
    a, b = p
    nei = []
    if a+b > 40:
        return []
    if a > b:
        nei.append((a+1, b))
        nei.append((a+2, b))
        nei.append((a+3, b))
        nei.append((a, b*2))
    if b > a:
        nei.append((a, b+1))
        nei.append((a, b+2))
        nei.append((a, b+3))
        nei.append((a*2, b))
    if a == b:
        nei.append((a, b+1))
        nei.append((a, b+2))
        nei.append((a, b+3))
        nei.append((a+1, b))
        nei.append((a+2, b))
        nei.append((a+3, b))
    return nei

import functools
@functools.cache
def coloring(p):
    colors = [coloring(x) for x in neighbours(p)]
    trues = [dist for x, dist in colors if not x]
    falses = [dist for x, dist in colors if x]
    if len(trues) > 0:
        return True, min(trues) + 1
    if len(falses) > 0:
        return False, max(falses)
    return False, 0

# 21
for s in range(1, 23):
    pos = (17, s)
    ifWinner, dist = coloring(pos)
    if (not ifWinner) and dist == 2:
        print(s)

# 20
for s in range(1, 36):
    pos = (5, s)
    ifWinner, dist = coloring(pos)
    if ifWinner and dist==2:
        print(s)

# 19
for s in range(1, 100):
    for a in range(1, s):
        b = s-a
        ifWinner, dist = coloring((a,b))
        if ifWinner:
            if dist == 1:
                print(coloring((a,b)))
                print(s)
                exit()
