import random


def merge(list, left, right):
    if left < right:
        mid = (left + right) // 2
        merge(list, left, mid)
        merge(list, mid + 1, right)

        tmp = []
        i = left
        j = mid + 1

        while i <= mid and j <= right:
            if list[i] <= list[j]:
                tmp.append(list[i])
                i += 1
            else:
                tmp.append(list[j])
                j += 1
        if i <= mid:
            tmp.extend(list[i:mid + 1])

        if j <= right:
            tmp.extend(list[j:right + 1])

        list[left:right + 1] = tmp


list = []
for i in range(1, 51):
    list.append(i)

random.shuffle(list)
merge(list, 0, len(list) - 1)
print(list)
