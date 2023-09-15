import json
import random

def load_questions(file_path="data/questions.json"):
    with open(file_path, "r", encoding="utf-8") as file:
        questions = json.load(file)
    return questions

def load_tags(file_path="data/tags.json"):
    with open(file_path, "r", encoding="utf-8") as file:
        tags = json.load(file)
    return tags

def save_questions(questions, file_path="data/questions.json"):
    with open(file_path, "w", encoding="utf-8") as file:
        json.dump(questions, file, ensure_ascii=False, indent=2)

def save_tags(tags, file_path="data/tags.json"):
    with open(file_path, "w", encoding="utf-8") as file:
        json.dump(tags, file, ensure_ascii=False, indent=2)

def generate_question_id(questions):
    return max([question["question_id"] for question in questions]) + 1

def generate_tag_id(tags):
    return max([tag["tag_id"] for tag in tags]) + 1

def generate_random_question(questions):
    return random.choice(questions)

def generate_random_tag(tags):
    return random.choice(tags)