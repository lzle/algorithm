import random

def bubble(list):
    nums = len(list)

    for i in range(nums-1):
        exchanged = False
        for j in range(nums-i- 1):
            if list[j] > list[j+1]:
                list[j], list[j+1] = list[j+1], list[j]
                exchanged = True

        if not exchanged:
            return

list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
bubble(list)
print(list)
