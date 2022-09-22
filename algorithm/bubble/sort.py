import random

def bubble(list):
    num = len(list)
    for i in range(num - 1):
        for j in range(1, num - i):
            if list[j] < list[j-1]:
                list[j - 1], list[j] = list[j], list[j - 1]

list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
bubble(list)
print(list)
