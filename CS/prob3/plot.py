
# ==============================================================================
# Python Import Session
# ==============================================================================
import matplotlib.pyplot as plt
import seaborn
# ==============================================================================
# Receive Sequence
# ==============================================================================

#temp = open("Data/test_taylor.csv")
temp = open("Data/taylor.csv")
List = [line.split(',') for line in temp]
X, Y, Z, T = [], [], [], []
for i, group in enumerate(List):
    X.append(float(group[0]))
    Y.append(float(group[1]))
    Z.append(float(group[2]))
    T.append(i / 2)
# ==============================================================================
# Plotting
# ==============================================================================
plt.figure(figsize=(10,10), dpi=600)
plt.plot(T, X, label='x Coord')
plt.plot(T, Y, label='y Coord')
plt.plot(T, Z, label='z Coord')
plt.title("Orbit")
plt.xlabel('Time')
plt.ylabel('Coordinate')
plt.legend()
plt.savefig('Fig/Orbit_Taylor.png')
