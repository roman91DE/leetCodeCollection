class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        
        ret = ""
        arr = sorted(strs, key=lambda x: x.__len__())
        shortest = arr[0]

        if len(shortest) == 0:
            return ret

        for i in range(len(shortest)):
            prefix = shortest[:i+1]

            for elem in arr[1:]:
                if elem[:i+1] != prefix:
                    return prefix[:-1]

        return prefix


        