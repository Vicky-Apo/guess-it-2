## Guess It 2

## 📌 Overview
This program predicts a **range for the next number** in a sequence using an **Adaptive Smoothing Algorithm**. It dynamically adjusts predictions based on **recent changes in the data**, ensuring a balance between **responsiveness to fluctuations** and **stability** within a dynamically adjusted range.

## 🔹 How It Works
### **1️⃣ Adaptive Trend Adjustment**
- The program maintains a **trend estimate** (`trendPoint`), which adjusts based on incoming values.
- The adjustment speed (`changeSpeed`) depends on how much the last value changed:
  - **Small changes** → Slow adjustment (10%)
  - **Moderate changes** → Medium adjustment (15%)
  - **Large changes** → Fast adjustment (25%)

### **2️⃣ Dynamic Range Prediction**
- Instead of a **fixed ±6 range**, the prediction **expands if recent jumps were large**.
- **Formula for range:**
  ```
  Lower Bound = trendPoint - dynamicRange
  Upper Bound = trendPoint + dynamicRange
  ```
  - `dynamicRange = 6.0 + (recentChange * 0.2)`

### **3️⃣ Continuous Learning**
- The program continuously refines its predictions **with each new value**.
- If the first number is given, it prints:
  ```
  No prediction available
  ```
  since there’s no previous data to compare.

---

## 📌 How to Run the Program
### **1️⃣ Ensure the Script Has Execution Permissions**
Run the following command:
```sh
chmod +x script.sh
```

### **2️⃣ Run the Test Environment**
Download the [tester zip file](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip) and place the **student/** folder in the root directory.

Start the environment with:
```sh
docker compose up --build
```
Then open **[http://localhost:3000](http://localhost:3000)** in your browser.

### **3️⃣ Add a Guesser (If Needed)**
If required, modify the URL:
```
http://localhost:3000?guesser=big-range
```

### **4️⃣ Run the Required Tests**
- You must pass **Test4** and **Test5** with the following guessers:
  - `big-range`
  - `linear-regr`
  - `correlation-coef`
  - `mse`
  - `nic`

Each test must be run **3 times**, and you need to pass **at least 2 out of 3** to pass the audit.

---

## 📌 Code Breakdown
### **🔹 Key Functions**
#### **1️⃣ adjustPrediction(newValue, lastValue)**
- Adjusts the trend estimate (`trendPoint`) based on the new value.
- Uses **adaptive smoothing** to avoid overcorrection.

#### **2️⃣ Dynamic Range Calculation**
- Expands the range dynamically based on **recent changes**:
  ```go
  dynamicRange := 6.0 + (math.Abs(newValue - lastValue) * 0.2)
  ```

#### **3️⃣ First Value Handling**
- If no previous data exists, prints:
  ```
  No prediction available
  ```

---

## 🔹 Why This Approach?
✅ **Balances accuracy and stability** by dynamically adjusting the trend.  
✅ **More efficient than storing all past values**, since it only needs the last number.  
✅ **Adaptable to different data patterns**, making it **more effective than a fixed range.**  

---

## 📌 Troubleshooting
### **1️⃣ Script Permission Denied**
If you see:
```
/bin/sh: 1: ./student/script.sh: Permission denied
```
Run:
```sh
chmod +x student/script.sh
```

### **2️⃣ Docker Issues**
If Docker doesn’t start:
```sh
docker compose down
docker compose up --build
```


## ✨ Final Notes
This project **successfully predicts number ranges using an adaptive trend model**. The approach ensures **fast reaction to changes** while preventing **overcorrection**.  

---

### **🔹 Authors**
- _Vicky Apostolou_

