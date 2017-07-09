import matplotlib.pyplot as plt
import seaborn

temp = open("Data/Circ.csv")
List = [line.split(',') for line in temp]
X, Y, Z = [], [], []
for group in List:
    X.append(float(group[0]))
    Y.append(float(group[1]))
    Z.append(float(group[2]))

plt.figure(dpi=600)
plt.plot(X, Y)
plt.plot(X, Z)
plt.title("Circ Plot")
plt.xlabel('X')
plt.ylabel('Y')
plt.savefig('Fig/Circ.png')