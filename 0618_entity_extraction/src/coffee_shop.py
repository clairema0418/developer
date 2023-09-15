import json
import random

class CoffeeShop:
    def __init__(self, questions_file="data/questions.json"):
        self.questions_file = questions_file
        self.questions = self.load_questions()

    def load_questions(self):
        with open(self.questions_file, "r", encoding="utf-8") as file:
            questions = json.load(file)
        return questions

    def get_random_question(self):
        return random.choice(self.questions)

    def generate_questions(self, num_questions=50):
        generated_questions = []
        for _ in range(num_questions):
            question = self.get_random_question()
            generated_questions.append(question)
        return generated_questions

    def save_generated_questions(self, generated_questions, output_file="data/generated_questions.json"):
        with open(output_file, "w", encoding="utf-8") as file:
            json.dump(generated_questions, file, ensure_ascii=False, indent=2)