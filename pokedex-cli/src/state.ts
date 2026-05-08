import { createInterface, type Interface } from "node:readline";
import { stdin, stdout } from "node:process"
import { commandExit } from "./command_exit.js";
import { commandHelp } from "./command_help.js"
import { PokeAPI } from "./pokeapi.js";
import { commandMap, commandMapBack } from "./command_map.js";


export type CLICommand = {
  name: string;
  description: string;
  callback: (state: State) => Promise<void>;
}

export type State = {
  rl: Interface;
  commands: Record<string, CLICommand>;
  pokeApi: PokeAPI;
  nextLocationsURL: string | null;
  prevLocationsURL: string | null;
}

export function initState(): State {
  const commands = {
    help: {
      name: "help",
      description: "Displays a help message",
      callback: commandHelp,
    },
    map: {
      name: "map",
      "description": "Get the next 20 Location Areas",
      callback: commandMap,
    },
    mapb: {
      name: "mapb",
      "description": "Get the previous 20 Location Areas",
      callback: commandMapBack,
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
    pokeApi: new PokeAPI(),
    nextLocationsURL: null,
    prevLocationsURL: null,
  };

  return initState
}
