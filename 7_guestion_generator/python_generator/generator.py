import os
import sys
import json
import time
import random
import requests

def generate_questions(question_count):
    questions = []
    for _ in range(question_count):
        question = {
            "id": random.randint(1, 1000000),
            "content": f"Question {random.randint(1, 1000000)}",
            "created_at": time.strftime("%Y-%m-%d %H:%M:%S"),
            "updated_at": time.strftime("%Y-%m-%d %H:%M:%S")
        }
        questions.append(question)
    return questions

def call_go_script(questions):
    url = "http://localhost:8080/api/question-generation-completed"
    headers = {"Content-Type": "application/json"}
    data = json.dumps({"questions": questions})
    response = requests.post(url, headers=headers, data=data)
    if response.status_code == 200:
        print("Go script notified successfully")
    else:
        print("Failed to notify Go script")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python generator.py <question_count>")
        sys.exit(1)

    question_count = int(sys.argv[1])
    questions = generate_questions(question_count)
    call_go_script(questions)