import { State } from "./state.js";

export function commandHelp(state: State) {
  const commands = state.commands;
  const start = `Welcome to the Pokedex!
Usage:
`
  console.log(start)
  Object.entries(commands).forEach(([_, cliCommand]) => {
    console.log(`${cliCommand.name}: ${cliCommand.description}`)
  })
}
