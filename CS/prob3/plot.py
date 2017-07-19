# ==============================================================================
# Python Import Session
# ==============================================================================

import matplotlib.pyplot as plt
import seaborn

# ==============================================================================
# Receive Sequence
# ==============================================================================

#temp = open("Data/test_taylor.csv")
temp1 = open("Data/X_ver.csv")
temp2 = open("Data/Energy.csv")
List1 = [line.split(',') for line in temp1]
List2 = [line.split(',') for line in temp2]
X, T, E = [], [], []
for group in List1:
    X.append(float(group[0]))
    T.append(float(group[1]))
for group in List2:
    E.append(float(group[0]))
# ==============================================================================
# Plotting
# ==============================================================================
plt.figure(figsize=(10,10), dpi=600)
plt.plot(T, X, label='x Coord')
plt.title("Orbit")
plt.xlabel('Time')
plt.ylabel('X')
plt.legend()
plt.savefig('Fig/X_Verlet.png')


plt.figure(figsize=(10,10), dpi=600)
plt.plot(T, E, label='Energy')
plt.title("Orbit")
plt.xlabel('Time')
plt.ylabel('E')
plt.legend()
plt.savefig('Fig/E.png')