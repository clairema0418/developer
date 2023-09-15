import json
from coffee_shop import CoffeeShop
from question_tagger import QuestionTagger

def main():
    coffee_shop = CoffeeShop()
    question_tagger = QuestionTagger()

    with open("data/questions.json", "r", encoding="utf-8") as file:
        questions = json.load(file)

    tagged_questions = []
    for question in questions:
        question_text = question["prompt"]
        tag = question_tagger.tag_question(question_text)
        tagged_questions.append({
            "question_id": question["question_id"],
            "question_text": question_text,
            "tag_id": tag["tag_id"],
            "tag_text": tag["tag_text"]
        })

    with open("data/tags.json", "w", encoding="utf-8") as file:
        json.dump(tagged_questions, file, ensure_ascii=False, indent=2)

if __name__ == "__main__":
    main()