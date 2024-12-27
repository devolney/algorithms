import numpy as np
import matplotlib.pyplot as plt

# Set the range of values ​​to n
n = np.arange(1, 6)

# Define the complexity functions
O_1 = np.ones_like(n)  # O(1)
O_n = n               # O(n)
O_n2 = n**2           # O(n^2)
O_logn = np.log(n)    # O(log n) (natural logaritm)
O_2n = 2**n           # O(2^n)

# Create the graph
plt.figure(figsize=(8, 6))

# Plot the functions
plt.plot(n, O_1, label="O(1)", linestyle='-', color='blue')
plt.plot(n, O_n, label="O(n)", linestyle='-', color='green')
plt.plot(n, O_n2, label="O(n²)", linestyle='-', color='red')
plt.plot(n, O_logn, label="O(log n)", linestyle='-', color='purple')
plt.plot(n, O_2n, label="O(2^n)", linestyle='-', color='orange')

# Add title and labels
plt.title("Big O curves comparison")
plt.xlabel("n")
plt.ylabel("f(n)")
plt.yscale('linear')
plt.xscale('linear')

# Add legend
plt.legend()

# Display the plot
plt.grid(True)
plt.show()

