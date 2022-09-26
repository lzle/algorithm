import random


def insertion(list):
    num = len(list)

    for i in range(1, num):
        curr = list[i]
        flag = i
        for j in range(i - 1, -1, -1):
            if curr >= list[j]:
                break
            list[j + 1] = list[j]
            flag = j

        if flag != i:
            list[flag] = curr


list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
insertion(list)
print(list)
