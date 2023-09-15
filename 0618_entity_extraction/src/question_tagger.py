import json
import random

class QuestionTagger:

    def __init__(self):
        self.questions = self.load_questions("data/questions.json")
        self.tags = self.load_tags("data/tags.json")

    def load_questions(self, file_path):
        with open(file_path, "r", encoding="utf-8") as file:
            questions = json.load(file)
        return questions

    def load_tags(self, file_path):
        with open(file_path, "r", encoding="utf-8") as file:
            tags = json.load(file)
        return tags

    def tag_question(self, question_id):
        question = self.questions[question_id]
        tagged_question = {
            "question_id": question_id,
            "question_text": question["question_text"],
            "tags": []
        }
        for tag_id, tag in self.tags.items():
            if tag["tag_text"] in question["question_text"]:
                tagged_question["tags"].append(tag_id)
        return tagged_question

    def generate_tagged_questions(self, num_cases):
        tagged_questions = []
        for _ in range(num_cases):
            question_id = random.choice(list(self.questions.keys()))
            tagged_question = self.tag_question(question_id)
            tagged_questions.append(tagged_question)
        return tagged_questions

    def save_tagged_questions(self, tagged_questions, file_path):
        with open(file_path, "w", encoding="utf-8") as file:
            json.dump(tagged_questions, file, ensure_ascii=False, indent=2)

if __name__ == "__main__":
    question_tagger = QuestionTagger()
    tagged_questions = question_tagger.generate_tagged_questions(50)
    question_tagger.save_tagged_questions(tagged_questions, "data/tagged_questions.json")