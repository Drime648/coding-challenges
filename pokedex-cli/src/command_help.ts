import { State } from "./state.js";

export async function commandHelp(state: State) {
  const commands = state.commands;
  const start = `Welcome to the Pokedex!
Usage:
`
  console.log(start)
  Object.entries(commands).forEach(([, cliCommand]) => {
    console.log(`${cliCommand.name}: ${cliCommand.description}`)
  })
}
