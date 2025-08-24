# 📦 Cointainers Packer 

A simple Go web application that demonstrates a **containers packing algorithm**.  
Users can enter available container sizes and a number of goods, and the server will calculate how to pack them efficiently.  

It also includes a `/tests` route that runs the Go test suite (compiled into a test binary) and shows results in the browser.  

---

## 🚀 Features

- **Web form** to input:
  - Container sizes (unique integers > 1)
  - Number of goods to pack
- **Packing algorithm** that finds the best way to fill goods into available containers
- **Bootstrap UI** with dynamic form powered by jQuery
- **Test runner** (`/tests` route) that executes project tests and displays results in browser
- Dockerized with a **multi-stage build** for minimal production image size

---

## 🧮 Algorithm (Packer) – Detailed Explanation

The goal of the algorithm is to **pack a given number of goods into containers** using the **fewest containers possible** while minimizing empty space in used containers.  

### Step 1 – Initialization

1. Start with a **list of container sizes** (e.g., `[250, 500, 1000, 2000]`).
2. Create an initial **queue of candidate packing approaches**, each containing a single container size.
3. Track:
   - `storedGoods` – total goods currently packed.
   - `containersTotal` – number of containers used.

---

### Step 2 – Iterative Expansion

1. For each approach in the queue:
   - Check if `storedGoods >= goodsSum` → mark as “completed.”
   - Otherwise, **try adding each container size** to this approach to create a **new approach**.

2. For each new approach:
   - Skip it if a better (or equal) approach with the same `storedGoods` already exists (pruning).
   - Otherwise, add it to the candidate queue.

---

### Step 3 – Selecting the Best Approach

1. After all new candidates are generated, **promote candidates** to the main queue.
2. Compare each approach to the **current best**:
   - First, prefer approaches with **smaller remainder**: `abs(goodsSum - storedGoods)`.
   - If the remainder is equal, prefer the approach with **fewer containers**.
   - If the remainder is bigger, but the approach with bigger remainder **is finished** and has **fewer containers**,
   while the other one **is not finished**, prefer the **finished** one
3. Repeat Step 2–3 until all approaches are “completed.”

---

### Step 4 – Result Construction

1. Once all approaches are completed, select the **best approach** according to the rules above.
2. Return a map of container size → count for the optimal packing solution.

---

### Step 5 – Example

**Input:**  
- Container sizes: `[250, 500, 1000, 2000]`  
- Goods: `4300`

**Algorithm flow:**  

1. Start candidates: `[250], [500], [1000], [2000]`  
2. Expand `[250]` → `[250, 250]`, `[250, 500]`, `[250, 1000]`, …  
3. Expand `[500]` → `[250, 500]` (won't be used, already have such entry), `[500, 500]`, …  
4. Continue until all approaches reach or exceed `4300` goods.  

**Output:**  

| Container | Count |
|-----------|-------|
| 500       | 1     |
| 2000      | 2     |

- Total containers: 3  
- Empty space left: 200

---

## 🐳 Running with Docker

#### ⚙️ Environment Variables

Set in `.env` (used by Docker Compose):

```
PORT=8080
```

Default: server runs on port `8080`.

---

### 1. Build and start containers

```bash
docker-compose up --build
```

### 2. Access the app

Replace **8080** port with you custom one, if you specified one in `.env` file

- Web app: [http://localhost:8080](http://localhost:8080)  
- Run tests: click the **"Run Tests"** button in UI or open [http://localhost:8080/tests](http://localhost:8080/tests)
