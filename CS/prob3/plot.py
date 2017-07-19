# ==============================================================================
# Python Import Session
# ==============================================================================

import matplotlib.pyplot as plt
import seaborn

# ==============================================================================
# Receive Sequence
# ==============================================================================

temp = open("Data/test.csv")
List = [line.split(',') for line in temp]
X, T = [], []
for group in List:
    X.append(float(group[0]))
    T.append(float(group[1]))

# ==============================================================================
# Plotting
# ==============================================================================

plt.figure(figsize=(10,10), dpi=600)
plt.plot(T, X, label='X Coord')
plt.title("Orbit")
plt.xlabel('Time')
plt.ylabel('X')
plt.legend()
plt.savefig('Fig/X.png')