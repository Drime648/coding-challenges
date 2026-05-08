import { getCommands } from "./command.js";

export function commandHelp() {
  const commands = getCommands();
  const start = `Welcome to the Pokedex!
Usage:
`
  console.log(start)
  Object.entries(commands).forEach(([_, cliCommand]) => {
    console.log(`${cliCommand.name}: ${cliCommand.description}`)
  })
}
