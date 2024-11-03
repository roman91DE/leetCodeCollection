class Solution:
    def rotateString(self, s: str, goal: str) -> bool:

        def shift(s: str) -> str:
            return s[1:] + s[0]

        if s == goal:
            return True

        elif (len(s) != len(goal)) or (set(s) != set(goal)):
            return False

        else:
            shifted = s

        for _ in range(len(s)):

            shifted = shift(shifted)

            if shifted == goal:
                return True

        return False

        