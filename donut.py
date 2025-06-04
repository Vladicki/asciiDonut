import math
import time
import os

# Donut configuration
R1 = 1        # Radius of tube
R2 = 2        # Distance from center to tube center
K2 = 5        # Distance from viewer
K1 = 15       # Scaling factor for projection

# Rotation angles
A = 0.0       # Rotation around X
B = 0.0       # Rotation around Z

# Screen dimensions
width, height = 80, 24

while True:
    # Clear screen buffer
    output = [' '] * (width * height)
    zbuffer = [0] * (width * height)

