def sortedSquaredArray(array):
    z = len(array)
    s = 0
    p = z - 1
    i = 0
    ans = []
    
    while i < z:
        ans.append(0)
        i += 1
    while z > 0:
        if abs(array[s]) >= abs(array[p]):
            ans[z - 1] = array[s] * array[s]
            s += 1
        else:
            ans[z - 1] = array[p] * array[p]
            p -= 1
        z-=1
    return ans