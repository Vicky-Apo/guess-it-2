🎯 Guess it 2

🔢 Predictive Range Guesser

📝 Overview

This program implements an Adaptive Smoothing Algorithm to predict a range for the next number in a sequence. The focus is on dynamically adjusting predictions based on recent changes in the data, ensuring a balance between responsiveness to fluctuations and maintaining stability within a dynamically adjusted range.

🔍 How It Works

📊 Core Concepts

⚙️ Adaptive Smoothing Algorithm

The algorithm adjusts predictions dynamically based on the magnitude of recent changes in the sequence.
The smoothing factor is not fixed but adapts based on the difference between the last two numbers:

Small changes → Lower adjustment rate (10%)

Moderate changes → Balanced adjustment rate (15%)

Large changes → Higher adjustment rate (25%)

📏 Prediction Ranges

The prediction is centered around an adjusted value that moves toward the new input by a percentage-based adjustment.
A dynamically adjusted range is applied to ensure flexibility in predictions.
Example: If the adjusted prediction is 150 and the recent change was significant, the output range will be broader than if the change was minimal.

🔄 Adaptation to Data

The model continuously updates based on incoming values, allowing it to adjust quickly to sudden spikes while remaining stable during steady sequences.

🌟 Key Features

📡 Dynamic Prediction

Predictions continuously refine as new values are read.

The algorithm ensures faster responses to sudden increases or decreases while preventing overcorrection.

🎯 Range Estimation

The dynamically adjusted range allows for better handling of both stable and volatile data.

The adaptive smoothing approach ensures more accurate predictions without being overly sensitive to outliers.

⚡ Efficiency

The program relies only on the previous value and its difference from the new value, making it computationally lightweight while still responsive to trends.

🔧 Configuration

Smoothing Factors:

Adjust the smoothing factors to fine-tune responsiveness based on changes in input values.

Example:

smallChangeFactor = 0.10
moderateChangeFactor = 0.15
largeChangeFactor = 0.25

Prediction Range:

The dynamically adjusted range ensures flexibility in predictions.

Example:

baseRange = 6.0
rangeAdjustmentFactor = 0.2

📚 Study

Adaptive Smoothing Algorithm

🛠️ Test the Program

Check if script.sh has executable permissions after cloning. Sometimes these permissions can be changed when cloned if the environment is different than the one when the file was pushed to the repository.
If the file has executable permissions, proceed to the next steps. Otherwise, give the permission using:

chmod +x script.sh

Run the test by downloading this zip file containing the tester. You should place the student/ folder in the root directory of the items provided.

Verify that Docker Desktop is running.

These commands should be run (on the root directory of files downloaded) to install dependencies and start the webpage on port 3000:

docker compose up --build

After opening your browser on the port 3000, if you try clicking on any of the Test Data buttons, you will notice that in the Dev Tools/ Console there is a message indicating that you need another guesser besides the student.

Adding a guesser is simple. You need to add in the URL a guesser, in other words, the name of one of the files present in the ai/ folder:

?guesser=<name_of_guesser>

For example:

?guesser=big-range

After that, choose one of the random datasets to test. You can wait for the program to test all of the values, or you can click Quick to skip the waiting and be presented with the results.

Since the website uses large datasets, we advise clearing the displays by clicking on the Clean button after each test.

You will need to test big-range, linear-regr, correlation-coef, mse, and nic.
For each of these, you will need to test Test4 and Test5 three times each.
If the student program wins at least 2/3 of each test, they pass the audit.

✍️ Authors
- _Vicky Apostolou_

