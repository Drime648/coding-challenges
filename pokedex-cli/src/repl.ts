import { stdin, stdout } from "node:process"
import { createInterface } from "node:readline"
import { getCommands } from "./command.js";



export function cleanInput(input: string): string[] {
  const trimmed = input.trim().toLowerCase();
  return trimmed ? trimmed.split(/\s+/) : []
}

export function startREPL() {
  const commands = getCommands();
  const repl = createInterface({
    input: stdin,
    output: stdout,
    prompt: "Pokedex > "
  });
  repl.prompt();
  repl.on("line", (input: string) => {
    const tokens = cleanInput(input)
    const command = tokens[0];
    if (command) {
      commands[command].callback();
    }
    repl.prompt();
  });
}

