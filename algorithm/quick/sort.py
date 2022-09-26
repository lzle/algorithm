import random


def quick(list, left, right):
    if left < right:
        p = partition(list, left, right)
        quick(list, left, p-1)
        quick(list, p+1, right)


def partition(list, left, right):
    pivot = list[right]
    i = left

    for j in range(left, right):
        if list[j] <= pivot:
            if i != j :
                list[j], list[i] = list[i], list[j]
            i += 1

    list[i], list[right] = list[right], list[i]
    return i


list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
quick(list, 0, len(list) - 1)
print(list)
