import json

def create_question(question, option_a, option_b, option_c, option_d, answer):
    return {
        "question": question,
        "option_a": option_a,
        "option_b": option_b,
        "option_c": option_c,
        "option_d": option_d,
        "answer": answer
    }

def main():
    questions = [
        create_question("台灣歷史上最早居住的族群是？", "荷蘭人", "西班牙人", "原住民族群", "漢族人口", "C"),
        create_question("根據《馬關條約》，哪一國家接管了台灣？", "荷蘭", "西班牙", "日本", "清朝", "C"),
        create_question("台灣在哪個世紀開始經歷殖民統治？", "16 世紀", "17 世紀", "18 世紀", "19 世紀", "B")
    ]

    with open("output/questions.json", "w", encoding="utf-8") as f:
        json.dump(questions, f, ensure_ascii=False, indent=2)

if __name__ == "__main__":
    main()