import matplotlib.pyplot as plt
import seaborn

temp = open("Data/test.csv")
List = [line.split(',') for line in temp]
X, Y = [], []
for group in List:
    X.append(float(group[0]))
    Y.append(float(group[1]))

plt.figure(dpi=600)
plt.plot(X, Y, label='y=sinx')
plt.title("Sin Plot")
plt.xlabel('X')
plt.ylabel('Y')
plt.legend()
plt.savefig('Fig/Test.png')