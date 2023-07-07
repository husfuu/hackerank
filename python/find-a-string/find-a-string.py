def count_substring(string, sub_string):
    num_string = len(sub_string)
    count = 0
    for i in range(len(string)):
        word = string[i:num_string+i]
        if word == sub_string:
            count += 1
    return count

if __name__ == '__main__':
    string = input().strip()
    sub_string = input().strip()
    
    count = count_substring(string, sub_string)
    print(count)