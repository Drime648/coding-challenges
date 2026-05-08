import { createInterface, type Interface } from "node:readline";
import { stdin, stdout } from "node:process"
import { commandExit } from "./command_exit.js";
import { commandHelp } from "./command_help.js"
export type CLICommand = {
  name: string;
  description: string;
  callback: (state: State) => void;
}

export type State = {
  rl: Interface;
  commands: Record<string, CLICommand>
}

export function initState(): State {
  const commands = {
    help: {
      name: "help",
      description: "Displays a help message",
      callback: commandHelp,
    },
    exit: {
      name: "exit",
      description: "Exits the pokedex",
      callback: commandExit,
    }
  };

  const repl = createInterface({
    input: stdin,
    output: stdout,
    prompt: "Pokedex > "
  });

  const initState: State = {
    rl: repl,
    commands: commands,
  };

  return initState
}
