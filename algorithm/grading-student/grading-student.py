def getMult5Data():
    # get all multiple by 5 data from 0-100
    mult_5_list = []
    for i in range(100):
        if ((i+1) % 5 == 0):
            mult_5_list.append(i+1)
    return mult_5_list


def gradingStudents(grades):
    for i in range(len(grades)):
        if grades[i] < 38:
            continue
        if (grades[i]+1) % 5 == 0:
            grades[i] += 1
        elif (grades[i]+2) % 5 == 0:
            grades[i] += 2
    return grades

# solusi pertama gw yg sangat tidak optimize (on^2)
# def gradingStudents(grades):
#     mult_5_list = getMult5Data()
#     for grade in grades:
#         if grade < 40:
#             print(grade)
#             continue
#             # return grades
#         temp = 0
#         for i in mult_5_list:
#             if i > grade:
#                 temp = i
#                 break
#         diff = abs(temp-grade)
#         if diff < 3:
#             print (grade+diff)
#         else:
#             print(grade)
#         # return grade


values = [73, 67, 38, 33]

print(gradingStudents(values))