
# ==============================================================================
# Python Import Session
# ==============================================================================
import matplotlib.pyplot as plt
import seaborn
# ==============================================================================
# Receive Sequence
# ==============================================================================

temp = open("Data/Cubic.csv")
List = [line.split(',') for line in temp]
X, Y = [], []
for i, elem in enumerate(List):
    X.append(elem[0])
    Y.append(elem[1])
# ==============================================================================
# Plotting
# ==============================================================================
plt.figure(figsize=(10,10), dpi=600)
plt.plot(X, Y, label='Spline')
plt.title("Cubic")
plt.xlabel('x')
plt.ylabel('y')
plt.legend()
plt.savefig('Fig/Cubic.png')