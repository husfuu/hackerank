import unittest

import palindrome

class TestPalindrome(unittest.TestCase):
    def testPalindrome(self):
        result = palindrome("aba")
        self.assertTrue(result, "ini palindrome")


if __name__ == '__main__':
    unittest.main()