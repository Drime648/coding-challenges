import { useEffect, useState, useCallback } from "react";
import words from "./wordList.json";
import { HangmanDrawing } from "./HangmanDrawing.tsx";
import { HangmanWord } from "./HangmanWord.tsx";
import { Keyboard } from "./Keyboard.tsx";

function getWord() {
  return words[Math.floor(Math.random() * words.length)]; // random number from 0-1, multiplied by length to go from 0-words.length
}

function App() {
  const [wordToGuess, setWordToGuess] = useState(() => getWord());
  const [guessedLetters, setGuessedLetters] = useState<string[]>([""]);
  const incorrectLetters = guessedLetters.filter(
    (letter) => !wordToGuess.includes(letter),
  );

  const isLost = incorrectLetters.length >= 6;
  const isWinner = wordToGuess
    .split("")
    .every((letter) => guessedLetters.includes(letter));

  const addGuessedLetter = useCallback(
    (letter: string) => {
      if (guessedLetters.includes(letter) || isLost || isWinner) return;
      setGuessedLetters((currentLetters) => [...currentLetters, letter]);
    },
    [guessedLetters, isLost, isWinner],
  );

  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      const key = e.key;
      if (!key.match(/^[a-z]$/)) return; //return if doesn't match a-z regex

      e.preventDefault();
      addGuessedLetter(key);
    };

    document.addEventListener("keypress", handler);
    return () => {
      document.removeEventListener("keypress", handler);
    };
  }, [guessedLetters]);

  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      const key = e.key;
      if (key !== "Enter") return;
      e.preventDefault();
      setWordToGuess(getWord());
      setGuessedLetters([]);
    };

    document.addEventListener("keypress", handler);
    return () => {
      document.removeEventListener("keypress", handler);
    };
  }, []);
  return (
    <div
      style={{
        maxWidth: "800px",
        display: "flex",
        flexDirection: "column",
        gap: "2rem",
        margin: "0 auto",
        alignItems: "center",
      }}
    >
      <div
        style={{
          fontSize: "2rem",
          textAlign: "center",
        }}
      >
        {isWinner ? "Winner!" : ""}
        {isLost ? "Nice Try! The word was " + wordToGuess : ""}
      </div>
      <HangmanDrawing numberOfGuesses={incorrectLetters.length} />
      <HangmanWord
        guessedLetters={guessedLetters}
        wordToGuess={wordToGuess}
        reveal={isLost}
      />
      <div style={{ alignSelf: "stretch" }}>
        <Keyboard
          disabled={isWinner || isLost}
          activeLetters={guessedLetters.filter((letter) =>
            wordToGuess.includes(letter),
          )}
          inactiveLetters={incorrectLetters}
          addGuessedLetter={addGuessedLetter}
        />
      </div>
    </div>
  );
}

export default App;
