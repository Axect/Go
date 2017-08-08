import matplotlib.pyplot as plt
import seaborn
import numpy as np

temp = open("Data/lag.csv")
List = [line.split(',') for line in temp]
X = np.empty(len(List))
Y = np.empty(len(List))
for i, group in enumerate(List):
    X[i] = group[0]
    Y[i] = group[1]

plt.figure(figsize=(10, 10), dpi=600)
plt.plot(X, Y)
plt.title("Lagrange Interpolation")
plt.savefig("Fig/Lag.png")