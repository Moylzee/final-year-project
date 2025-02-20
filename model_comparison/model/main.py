import os
import pandas as pd
import glob
from sklearn.model_selection import train_test_split
from sklearn.feature_extraction.text import CountVectorizer
from sklearn.naive_bayes import MultinomialNB
from sklearn.metrics import classification_report

# Define paths
positive_path = 'data/positive_data'
negative_path = 'data/negative_data'

# Read CSV files into separate dataframes
positive_files = glob.glob(os.path.join(positive_path, '*.csv'))
negative_files = glob.glob(os.path.join(negative_path, '*.csv'))

X_positive = []
y_positive = []

for file in positive_files:
    df = pd.read_csv(file)
    # Assuming we're using the first column as our feature
    X_positive.append(df.iloc[:, 0].values.tolist())

y_positive = [1] * len(X_positive)

X_negative = []
y_negative = []

for file in negative_files:
    df = pd.read_csv(file)
    X_negative.append(df.iloc[:, 0].values.tolist())
    y_negative = [0] * len(X_negative)

# Combine all data
X = X_positive + X_negative
y = y_positive + y_negative

# Flatten the lists of lists into individual strings
X = [' '.join(text) for text in X]

# Split into training and testing sets
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Vectorize text data
vectorizer = CountVectorizer()
X_train_vectorized = vectorizer.fit_transform(X_train)
X_test_vectorized = vectorizer.transform(X_test)

# Train the model
clf = MultinomialNB()
clf.fit(X_train_vectorized, y_train)

# Evaluate the model
y_pred = clf.predict(X_test_vectorized)
print(classification_report(y_test, y_pred))
