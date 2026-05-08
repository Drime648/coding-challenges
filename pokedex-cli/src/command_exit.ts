import process from "node:process";
import { State } from "./state.js";

export async function commandExit(state: State, ..._args: string[]) {
  state.rl.close();
  state.pokeApi.cache.stopReapLoop();
  console.log("Closing the Pokedex... Goodbye!");
  process.exit(0);
}
