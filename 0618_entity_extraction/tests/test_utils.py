import unittest
from src.utils import generate_questions, generate_tags, generate_question_tag_relation

class TestUtils(unittest.TestCase):

    def test_generate_questions(self):
        questions = generate_questions(50)
        self.assertEqual(len(questions), 50)
        for question in questions:
            self.assertIn('question_id', question)
            self.assertIn('question_text', question)

    def test_generate_tags(self):
        tags = generate_tags()
        self.assertGreater(len(tags), 0)
        for tag in tags:
            self.assertIn('tag_id', tag)
            self.assertIn('tag_text', tag)

    def test_generate_question_tag_relation(self):
        questions = generate_questions(50)
        tags = generate_tags()
        question_tag_relation = generate_question_tag_relation(questions, tags)
        self.assertEqual(len(question_tag_relation), len(questions))
        for relation in question_tag_relation:
            self.assertIn('question_id', relation)
            self.assertIn('tag_id', relation)

if __name__ == '__main__':
    unittest.main()