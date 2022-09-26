import random


def selection(list):
    num = len(list)

    for i in range(num - 1):
        min = i
        for j in range(i + 1, num):
            if list[j] < list[min]:
                min = j

        list[i], list[min] = list[min], list[i]


list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
selection(list)
print(list)
