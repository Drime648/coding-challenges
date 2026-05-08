import process from "node:process";
import { State } from "./state.js";

export function commandExit(state: State) {
  console.log("Closing the Pokedex... Goodbye!");
  process.exit(0);
}
