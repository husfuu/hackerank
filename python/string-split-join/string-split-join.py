def split_and_join(line):
    line_list = line.split(" ")
    line = "-".join(line_list)
    return line

if __name__ == '__main__':
    line = input()
    result = split_and_join(line)
    print(result)