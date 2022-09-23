import random


def selection(list):
    nums = len(list)

    for i in range(nums-1):
        min = i
        for j in range(i+1, nums):
            if list[j] < list[min]:
                min = j

        list[i], list[min] = list[min], list[i]

list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
selection(list)
print(list)