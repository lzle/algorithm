import random


def heap(list):
    n = len(list)
    build_heap(list, n)

    for i in range(n - 1):
        list[0], list[n - 1 - i] = list[n - 1 - i], list[0]
        heapify(list, n - 1 - i, 0)


def build_heap(list, n):
    for i in range(n // 2 - 1, -1, -1):
        heapify(list, n, i)


def heapify(list, n, i):
    while True:
        max = i
        if i * 2 + 1 <= n - 1 and list[max] < list[i * 2 + 1]:
            max = i * 2 + 1

        if i * 2 + 2 <= n - 1 and list[max] < list[i * 2 + 2]:
            max = i * 2 + 2

        if max == i:
            break

        list[i], list[max] = list[max], list[i]
        i = max


list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
heap(list)
print(list)
