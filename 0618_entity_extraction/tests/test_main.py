import unittest
from src.main import generate_questions_and_tags

class TestMain(unittest.TestCase):

    def test_generate_questions_and_tags(self):
        questions_and_tags = generate_questions_and_tags(50)
        self.assertEqual(len(questions_and_tags), 50)

        for item in questions_and_tags:
            self.assertIn("prompt", item)
            self.assertIn("completion", item)

            prompt = item["prompt"]
            completion = item["completion"]

            self.assertIsInstance(prompt, str)
            self.assertIsInstance(completion, str)

            self.assertNotEqual(prompt.strip(), "")
            self.assertNotEqual(completion.strip(), "")

if __name__ == "__main__":
    unittest.main()