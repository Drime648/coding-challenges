import { initState } from "./state.js";


export function cleanInput(input: string): string[] {
  const trimmed = input.trim().toLowerCase();
  return trimmed ? trimmed.split(/\s+/) : []
}

export function startREPL() {
  let state = initState();
  const commands = state.commands;
  const repl = state.rl;
  repl.prompt();
  repl.on("line", async (input) => {
    const tokens = cleanInput(input)
    const command = tokens[0];
    if (command) {
      if (Object.hasOwn(commands, command)) {
        await commands[command].callback(state, ...tokens.slice(1));
      } else {
        console.log(`Invalid command: ${command}`);
      }
    }
    repl.prompt();
  });
}

