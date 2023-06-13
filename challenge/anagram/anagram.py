def is_anagram(word1, word2):
    if sorted(word1) == sorted(word2):
        return True

    return False

anagram_word = [
    {
        "race": "care",
    },
    {
        "part": "trap",
    },
    {
        "heart": "earth",

    },
    {
        "knee":"keen"
    }
]

for item in anagram_word:
    for word1, word2 in item.items():
        print(is_anagram(word1, word2))

not_anagram_word = [
    {
        "racex": "care",
    },
    {
        "parst": "traap",
    },
    {
        "h1eart": "ea2rth",

    },
    {
        "3knee":"3keen" # ini true :>
    }
]


for item in not_anagram_word:
    for word1, word2 in item.items():
        print(is_anagram(word1, word2))