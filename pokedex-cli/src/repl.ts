import { stdin, stdout } from "node:process"
import { createInterface } from "node:readline"



export function cleanInput(input: string): string[] {
  const trimmed = input.trim().toLowerCase();
  return trimmed ? trimmed.split(/\s+/) : []
}

export function startREPL() {
  const repl = createInterface({
    input: stdin,
    output: stdout,
    prompt: "Pokedex > "
  });
  repl.prompt();
  repl.on("line", (input: string) => {
    const [command] = cleanInput(input)
    if (command) {
      console.log(`Your command was: ${command}`)
    }
    repl.prompt();
  });
}

