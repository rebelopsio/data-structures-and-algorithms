def twoNumberSum(array, targetSum):
    nums = {}
    for num in array:
        potential_ans = targetSum - num
        if potential_ans in nums:
            return [num, potential_ans]
        else:
            nums[num] = True
    return []