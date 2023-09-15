from setuptools import setup, find_packages

setup(
    name="coffee_shop_question_tagger",
    version="0.1",
    packages=find_packages(),
    install_requires=[
        "json",
    ],
    entry_points={
        "console_scripts": [
            "coffee_shop_question_tagger = src.main:main",
        ],
    },
)