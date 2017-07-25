
# ==============================================================================
# Python Import Session
# ==============================================================================
import matplotlib.pyplot as plt
import seaborn
# This is the 3D plotting toolkit
from mpl_toolkits.mplot3d import Axes3D
# ==============================================================================
# Receive Sequence
# ==============================================================================

temp1 = open("Data/taylor.csv")
temp2 = open("Data/reverse.csv")
temp3 = open("Data/energy.csv")
temp4 = open("Data/energy_rev.csv")
List1 = [line.split(',') for line in temp1]
List2 = [line.split(',') for line in temp2]
List3 = [line.split(',') for line in temp3]
List4 = [line.split(',') for line in temp4]
X1, Y1, Z1, X2, Y2, Z2, t, T1, T2, U1, U2, E1, E2 = [],[],[],[],[],[],[],[],[],[], [], [], []
for i, _ in enumerate(List1):
    X1.append(float(List1[i][0]))
    Y1.append(float(List1[i][1]))
    Z1.append(float(List1[i][2]))
    t.append(i / 2)
    X2.append(float(List2[i][0]))
    Y2.append(float(List2[i][1]))
    Z2.append(float(List2[i][2]))
    T1.append(float(List3[i][0]))
    U1.append(float(List3[i][1]))
    E1.append(T1[i] + U1[i])
    T2.append(float(List4[i][0]))
    U2.append(float(List4[i][1]))
    E2.append(T2[i] + U2[i])
# ==============================================================================
# Plotting
# ==============================================================================
plt.figure(figsize=(10,10), dpi=600)
plt.plot(t, X1, label='x Coord')
plt.plot(t, Y1, label='y Coord')
plt.plot(t, Z1, label='z Coord')
plt.title("Orbit")
plt.xlabel('Time (Days)')
plt.ylabel('Coordinate (AU)')
plt.legend()
plt.savefig('Fig/Orbit_Taylor.png')

fig = plt.figure(figsize=(10,10), dpi=600)
ax = plt.axes(projection='3d')
ax.plot(X1, Y1, Z1, label="Position")
plt.title("Orbit")
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Z')
plt.legend()
plt.savefig('Fig/Orbit_3D.png')

plt.figure(figsize=(10,10), dpi=600)
plt.plot(t, X2, label='x Coord')
plt.plot(t, Y2, label='y Coord')
plt.plot(t, Z2, label='z Coord')
plt.title("Reverse Orbit")
plt.xlabel('Time (Days)')
plt.ylabel('Coordinate (AU)')
plt.legend()
plt.savefig('Fig/Orbit_Reverse.png')

plt.figure(figsize=(10,10), dpi=600)
plt.plot(t, T1, label='E_k')
plt.plot(t, U1, label='E_p')
plt.plot(t, E1, label='E_tot')
plt.title("Orbit")
plt.xlabel('Time (Days)')
plt.ylabel('Energy')
plt.legend()
plt.savefig('Fig/Energy.png')

plt.figure(figsize=(10,10), dpi=600)
plt.plot(t, T2, label='E_k')
plt.plot(t, U2, label='E_p')
plt.plot(t, E2, label='E_tot')
plt.title("Reverse Orbit")
plt.xlabel('Time (Days)')
plt.ylabel('Energy')
plt.legend()
plt.savefig('Fig/Energy_Reverse.png')
