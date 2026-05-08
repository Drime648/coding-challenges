import { State } from "./state.js";


export async function commandCatch(state: State, ...args: string[]) {
  if (args.length !== 1) {
    console.log("Must specify a pokemon!");
  }
  const name = args[0];
  try {
    const pokemonData = await state.pokeApi.fetchPokemon(name);
    console.log(`Throwing a Pokeball at ${name}...`);
    if (Math.random() * 1000 > pokemonData.base_experience) {
      console.log(`${name} was caught!`);
      state.pokedex.set(name, pokemonData);
    } else {
      console.log(`${name} escaped!`);
    }
  } catch {
    console.log(`${name} is not a real Pokemon!`);
  }
}
