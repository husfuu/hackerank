def swap_case(s):
    new_str = ""
    for char in s:
        if char.islower():
            new_str += char.upper()
        else:
            new_str += char.lower()
    return new_str


word = "Hai"

res = swap_case(word)
print(res)