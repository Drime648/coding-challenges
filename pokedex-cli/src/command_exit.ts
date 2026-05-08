import process from "node:process";

export function commandExit() {
  console.log("Closing the Pokedex... Goodbye!");
  process.exit(0);
}
