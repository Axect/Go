# ==============================================================================
# Python Import Session
# ==============================================================================

import matplotlib.pyplot as plt
import seaborn

# ==============================================================================
# Receive Sequence
# ==============================================================================

#temp = open("Data/test_taylor.csv")
#temp = open("Data/X_ver.csv")
temp = open("Data/Energy.csv")
List = [line.split(',') for line in temp]
E, T = [], []
for group in List:
    E.append(float(group[0]))
    T.append(float(group[1]))

# ==============================================================================
# Plotting
# ==============================================================================

plt.figure(figsize=(10,10), dpi=600)
plt.plot(T, E, label='Energy')
plt.title("Orbit")
plt.xlabel('Time')
plt.ylabel('E')
plt.legend()
plt.savefig('Fig/E.png')