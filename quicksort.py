a = [1,2,3,4,5,6,7,8,9]
b = [4,5,6,1,3,2,8,6,9]

print(a)

def quicksort(array):
    if len(array) < 2:
        return array
    else:
        pivot = array[0]
        less = [i for i in array[1:] if i <= pivot]
        greater = [i for i in array[1:] if i > pivot]
        return quicksort(less) + [pivot] + quicksort(greater)

print(quicksort(b))