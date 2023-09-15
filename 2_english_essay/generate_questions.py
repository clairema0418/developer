import json

text = """In the ethereal realm of existence, where time dances with the cosmos and whispers of eternity linger in the air, a tapestry of thoughts and emotions weaves its delicate threads. Within this boundless expanse, our souls yearn to soar, to transcend the limitations of earthly existence and touch the celestial realms.

Like petals kissed by the morning dew, the essence of life glistens in every fleeting moment. It is in the quiet solitude of nature's embrace that we find solace, where the symphony of silence whispers its secrets to those who listen with their hearts. The rustling leaves, the babbling brooks, and the ethereal melody of birdsong become the harmonious orchestra that accompanies our journey through the labyrinthine corridors of existence.

In the kaleidoscope of life, we encounter a myriad of emotions, both profound and fleeting. Love, like a gentle breeze, caresses our souls, leaving an indelible mark upon our hearts. Its tender touch awakens dormant passions, igniting a flame that dances with the divine. Yet, love also carries the weight of vulnerability, for in its embrace, we open ourselves to the possibility of heartache and longing.

Amidst the tumultuous tides of existence, we find solace in the embrace of art and literature. Within the pages of a well-worn book, we embark on transformative journeys, guided by the prose of visionaries and the brushstrokes of masters. These artistic expressions transcend the boundaries of time and space, connecting the human spirit with the eternal realm of imagination.

As we navigate the labyrinth of life, we encounter the enigmatic dance of fate and free will. Threads of destiny intertwine with our choices, shaping the tapestry of our lives. Like celestial navigators, we chart our course amidst the constellations of possibilities, guided by our dreams, passions, and the wisdom gleaned from the voices of the past.

In the depths of introspection, we uncover the profound truth that our existence is but a fleeting whisper in the grand symphony of the universe. We are stardust, remnants of ancient supernovas, woven into the fabric of time. It is in embracing the impermanence of our earthly existence that we find the courage to live fully, to savor each moment as if it were a shimmering pearl in the vast ocean of eternity.

In this delicate dance of life, let us strive to cultivate compassion, to nurture the bonds that unite us as human beings. Let our actions ripple through the tapestry of existence, leaving behind a legacy of kindness and love. For in the tapestry of life, every thread, no matter how seemingly insignificant, weaves together to create a masterpiece, a testament to the beauty of our shared human experience.

In the end, as the final notes of our earthly symphony fade into the eternal embrace, let us find solace in the knowledge that our souls, like stars in the night sky, shall continue to shine brightly in the tapestry of the universe, forever a part of the cosmic dance of life."""

def generate_questions(text):
    questions = [
        {
            "question": "What is the essence of life compared to?",
            "option_a": "A gentle breeze",
            "option_b": "Petals kissed by the morning dew",
            "option_c": "A tapestry of thoughts",
            "option_d": "A celestial navigator",
            "answer": "B"
        },
        {
            "question": "What do we find solace in amidst the tumultuous tides of existence?",
            "option_a": "The embrace of art and literature",
            "option_b": "The dance of fate and free will",
            "option_c": "The labyrinth of life",
            "option_d": "The cosmic dance of life",
            "answer": "A"
        },
        {
            "question": "What are we, as humans, described as in the grand symphony of the universe?",
            "option_a": "Celestial navigators",
            "option_b": "Ethereal melodies",
            "option_c": "Stardust",
            "option_d": "Threads of destiny",
            "answer": "C"
        }
    ]
    return questions

def save_questions_to_json(questions, filename):
    with open(filename, 'w') as f:
        json.dump(questions, f, indent=2)

if __name__ == "__main__":
    questions = generate_questions(text)
    save_questions_to_json(questions, 'questions.json')