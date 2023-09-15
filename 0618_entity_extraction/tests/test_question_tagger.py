import unittest
from src.question_tagger import QuestionTagger

class TestQuestionTagger(unittest.TestCase):

    def setUp(self):
        self.question_tagger = QuestionTagger()

    def test_tag_question(self):
        test_cases = [
            {
                "question": "你可以幫我介紹什麼產品嗎?",
                "expected_tags": ["product", "introduction"]
            },
            {
                "question": "這個咖啡多少錢?",
                "expected_tags": ["price", "coffee"]
            },
            {
                "question": "你們有無糖的選擇嗎?",
                "expected_tags": ["sugar-free", "options"]
            }
        ]

        for case in test_cases:
            result = self.question_tagger.tag_question(case["question"])
            self.assertEqual(set(result), set(case["expected_tags"]))

    def test_generate_question_tag_data(self):
        question_data = self.question_tagger.generate_question_tag_data(50)
        self.assertEqual(len(question_data), 50)

        for data in question_data:
            self.assertIn("question_id", data)
            self.assertIn("question_text", data)
            self.assertIn("tags", data)
            self.assertIsInstance(data["tags"], list)

if __name__ == '__main__':
    unittest.main()