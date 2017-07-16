# ==============================================================================
# Python Import Session
# ==============================================================================

import matplotlib.pyplot as plt
import seaborn

# ==============================================================================
# Receive Sequence
# ==============================================================================

temp = open("Data/physics.csv")
List = [line.split(',') for line in temp]
Name, Score, Rank = [], [], []
for group in List:
    Name.append(group[0])
    Score.append(int(group[1]))
    Rank.append(int(group[2]))
# ==============================================================================
# Plotting
# ==============================================================================

plt.figure(figsize=(10,10), dpi=600)
plt.plot(Rank, Score, label='Score')
plt.title("Distribution of Score")
plt.xlabel('Rank')
plt.ylabel('Score')
plt.legend()
plt.savefig('Fig/Distrib.png')