def palindrome(word):
    """ algorithm
    1. reverse the string using backward loop and store the result in empty string
    2. compare the reverse result with the given word
    """
    temp = ""
    # backward loop
    for i in range(len(word)-1, -1, -1):
        temp += word[i]

    if temp == word:
        return True
    return False


palindrome_data = ["aba", "a", "kodok"]
not_palindrome_data = ["ab", "abab", "kodcok"]

# test true palindrome
for data in palindrome_data:
    print(palindrome(data))

# test false palindrome
for data in not_palindrome_data:
    print(palindrome(data))
