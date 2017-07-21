import matplotlib.pyplot as plt
import seaborn as sns
import numpy as np

temp = open("Data/gauge.csv", 'r')
List = [line.split(',') for line in temp]
g1, g2, g3 = np.empty(len(List)), np.empty(len(List)), np.empty(len(List))

for i, group in enumerate(List):
    g1[i] = group[0]
    g2[i] = group[1]
    g3[i] = group[2]

x = np.array([i/10000 for i in range(len(g1))])

plt.figure(figsize=(8, 6), dpi=300)
plt.plot(x, g1, label="g1")
plt.plot(x, g2, label='g2')
plt.plot(x, g3, label="g3")
plt.legend()
plt.title("Running Gauge")
plt.xlabel('t')
plt.ylabel('gauge')
plt.savefig("Fig/Gauge.png")
