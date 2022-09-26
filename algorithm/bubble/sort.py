import random


def bubble(list):
    num = len(list)

    for i in range(num - 1):
        exchanged = False
        for j in range(num - i - 1):
            if list[j] > list[j + 1]:
                list[j], list[j + 1] = list[j + 1], list[j]
                exchanged = True

        if not exchanged:
            return


list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
bubble(list)
print(list)
