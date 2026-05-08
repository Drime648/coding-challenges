import { State } from "./state.js";



export async function commandInspect(state: State, ...args: string[]) {
  if (args.length !== 1) {
    console.log("Must specify a pokemon to inspect!");
  }
  const name = args[0];
  if (!state.pokedex.has(name)) {
    console.log("you have not caught that pokemon");
  } else {
    const pokemonData = state.pokedex.get(name);
    console.log(`Name: ${pokemonData?.name}`);
    console.log(`Height: ${pokemonData?.height}`);
    console.log(`Weight: ${pokemonData?.weight}`);
  }

}
