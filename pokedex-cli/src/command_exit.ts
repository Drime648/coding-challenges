import process from "node:process";
import { State } from "./state.js";

export async function commandExit(state: State) {
  state.rl.close();
  console.log("Closing the Pokedex... Goodbye!");
  process.exit(0);
}
