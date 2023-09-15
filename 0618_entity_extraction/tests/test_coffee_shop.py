import unittest
from src.coffee_shop import CoffeeShop

class TestCoffeeShop(unittest.TestCase):

    def setUp(self):
        self.coffee_shop = CoffeeShop()

    def test_generate_questions(self):
        questions = self.coffee_shop.generate_questions(50)
        self.assertEqual(len(questions), 50)

        for question in questions:
            self.assertIn("prompt", question)
            self.assertIn("completion", question)

    def test_export_to_json(self):
        questions = self.coffee_shop.generate_questions(50)
        json_output = self.coffee_shop.export_to_json(questions)
        self.assertIsInstance(json_output, str)

        try:
            import json
            json.loads(json_output)
        except ValueError:
            self.fail("export_to_json did not return valid JSON")

if __name__ == '__main__':
    unittest.main()