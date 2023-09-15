# Coffee Shop Question Tagger

This application simulates user questions in a coffee shop and tags them. It generates 50 examples in JSON format.

## Installation

1. Clone the repository.
2. Install the required packages using `pip install -r requirements.txt`.

## Usage

1. Run `src/main.py` to generate the questions and tags in JSON format.
2. The generated questions and tags will be saved in `data/questions.json` and `data/tags.json` respectively.

## Testing

To run the tests, execute the following command:

```
python -m unittest discover tests
```

## Files

- `data/questions.json`: Contains the generated questions.
- `data/tags.json`: Contains the generated tags.
- `src/main.py`: Main script to generate questions and tags.
- `src/coffee_shop.py`: Contains the CoffeeShop class to simulate user questions.
- `src/question_tagger.py`: Contains the QuestionTagger class to tag the questions.
- `src/utils.py`: Contains utility functions.
- `tests/test_main.py`: Test cases for main.py.
- `tests/test_coffee_shop.py`: Test cases for coffee_shop.py.
- `tests/test_question_tagger.py`: Test cases for question_tagger.py.
- `tests/test_utils.py`: Test cases for utils.py.
- `requirements.txt`: Required packages for the application.
- `setup.py`: Setup script for the application.
- `.gitignore`: Git ignore file.