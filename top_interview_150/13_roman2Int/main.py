#!/usr/bin/env python
# coding: utf-8

# In[57]:


class Solution:

    def romanToInt(self, s: str) -> int:

        d = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

        def recsum(s: str, acc: int) -> int:
            if len(s) == 0:
                return acc
            elif len(s) == 1:
                return acc + d[s]
            else:
                i, j = 0, 1
                if d[s[i]] < d[s[j]]:
                    acc += d[s[j]] - d[s[i]]
                    return recsum(s[j + 1 :], acc)
                else:
                    acc += d[s[i]]
                    return recsum(s[i + 1 :], acc)

        return recsum(s, 0)


s = Solution()


# In[60]:


assert s.romanToInt("VII") == 7


# In[61]:


assert s.romanToInt("IX") == 9
