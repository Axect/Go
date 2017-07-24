import matplotlib.pyplot as plt
import seaborn as sns
import numpy as np

temp = open("Data/Gauge1.csv", 'r')
List = [line.split(',') for line in temp]
l = len(List)
lH, g1, g2, g3, yt, phi, G, V = np.empty(l), np.empty(l), np.empty(l), np.empty(l), np.empty(l), np.empty(l), np.empty(l), np.empty(l)

for i, group in enumerate(List):
    lH[i] = group[0]
    g1[i] = group[1]
    g2[i] = group[2]
    g3[i] = group[3]
    yt[i] = group[4]
    phi[i] = group[5]
    G[i] = group[6]
    V[i] = group[7]

x = np.array([i/10000 for i in range(len(g1))])

plt.figure(figsize=(8, 6), dpi=300)
plt.plot(x, lH, label="lambda")
plt.plot(x, g1, label="g1")
plt.plot(x, g2, label='g2')
plt.plot(x, g3, label="g3")
plt.plot(x, yt, label="yt")
plt.legend()
plt.title("Running Gauge")
plt.xlabel('t=ln(phi/MpR)')
plt.ylabel('gauge')
plt.savefig("Fig/Gauge1.png")


plt.figure(figsize=(8, 6), dpi=300)
plt.plot(x, G, label="G")
plt.legend()
plt.title("Running G")
plt.xlabel('t=ln(phi/MpR)')
plt.ylabel('G')
plt.savefig("Fig/G1.png")

plt.figure(figsize=(8, 6), dpi=300)
plt.plot(phi, V, label="Potential")
plt.legend()
plt.title("Potential")
plt.xlabel('phi')
plt.ylabel('V')
plt.axis([0, 1, 0, 1e-06])
plt.savefig("Fig/Pot1.png")
