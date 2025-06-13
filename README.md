# 🍩 ASCII Spinning Donut in Go

This is a fun and mesmerizing ASCII art animation of a **3D spinning donut**, implemented in Go. It renders a rotating torus (donut shape) entirely with characters in your terminal, using basic 3D math, perspective projection, and luminance-based shading.

![donut](https://user-images.githubusercontent.com/66507909/135775594-dd729ae2-22d7-4692-92e9-c174c29991b2.gif)
---

## ✨ Features

- Real-time 3D rotation using Y-axis rotation only
- Smooth shading using ASCII luminance ramp
- Terminal output using Go standard library
- Super minimal and cross-platform

---

## 🧠 How it Works

- A 3D torus is generated using parametric equations
- Torus is rotated around the **XYZ-axes**
- Each point is projected into 2D using perspective projection
- A z-buffer is used to handle overlapping depth
- Luminance is calculated and mapped to a character for shading

---

## ▶️ Run It

### Prerequisites

- Go 1.18+ installed

### Build & Run

```bash
git clone https://github.com/Vladicki/asciiDonut.git
cd asciiDonut
go run ascii-donut.go

