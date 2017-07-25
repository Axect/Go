#import matplotlib as mpl
#mpl.use('Agg')
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns

MpR      = 2.4 * 1E+18

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
    phi[i] = float(group[5])/(MpR)
    G[i] = group[6]
    V[i] = float(group[7])/(MpR**4)

x = np.array([i/10000 for i in range(len(g1))])

a = input("Choose 1.Gauge, 2.Lambda, 3.G(t), 4.V(phi) :")
b = a.split(' ')
if '1' in b:
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

if '2' in b:
    plt.figure(figsize=(8, 6), dpi=300)
    plt.plot(x, lH, label="lH")
    plt.legend()
    plt.title("Running lH")
    plt.xlabel('t=ln(phi/MpR)')
    plt.ylabel('lH')
    plt.savefig("Fig/lH1.png")

if '3' in b:
    plt.figure(figsize=(8, 6), dpi=300)
    plt.plot(x, G, label="G")
    plt.legend()
    plt.title("Running G")
    plt.xlabel('t=ln(phi/MpR)')
    plt.ylabel('G')
    plt.savefig("Fig/G1.png")

if '4' in b:
    plt.figure(figsize=(8, 6), dpi=300)
    plt.plot(phi, V, label="Potential")
    plt.legend()
    plt.title("Potential")
    plt.xlabel('phi')
    plt.ylabel('V')
    for ph, v in zip(phi, V):
        if abs(ph - 1) < 1e-04:
            mv = v
            break
    plt.axis([0, 1, 0, mv+mv/10])
    plt.savefig("Fig/Pot1.png")
