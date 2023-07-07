if __name__ == '__main__':
    s = input()
    print(True in [c.isalnum() for c in s])
    print(True in [c.isalpha() for c in s])
    print(True in [c.isdigit() for c in s])
    print(True in [c.islower() for c in s])
    print(True in [c.isupper() for c in s])