import random


def selection(list):
    num = len(list)

    for i in range(num - 1):
        min = i

        for j in range(1, num - i):
            if list[j + i] < list[min]:
                min = j + i

        list[i], list[min] = list[min], list[i]

list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
selection(list)
print(list)