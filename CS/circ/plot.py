# ==============================================================================
# Python Import Session
# ==============================================================================

import matplotlib.pyplot as plt
import seaborn

# ==============================================================================
# Receive Sequence
# ==============================================================================

temp = open("Data/Circ.csv")
List = [line.split(',') for line in temp]
X, Y, Z = [], [], []
for group in List:
    X.append(float(group[0]))
    Y.append(float(group[1]))
    Z.append(float(group[2]))

temp2 = open("Data/Tan.csv")
List2 = [line.split(',') for line in temp2]
P, Q = [], []
for group in List2:
    P.append(float(group[0]))
    Q.append(float(group[1]))

# ==============================================================================
# Plotting
# ==============================================================================

plt.figure(figsize=(10,10), dpi=600)
plt.plot(X, Y, label='Circle')
plt.plot(X, Z)
plt.plot(P, Q, label='Tangent')
plt.title("Circle Plot")
plt.xlabel('X')
plt.ylabel('Y')
plt.legend()
plt.savefig('Fig/CircTan.png')