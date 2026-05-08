import { State } from "./state.js"

export async function commandExplore(state: State, ...args: string[]) {
  if (args.length !== 1) {
    console.log("Must specify location area!")
  }
  const locationAreaName = args[0];

  try {
    const locationArea = await state.pokeApi.fetchLocationArea(locationAreaName);
    locationArea.pokemon_encounters.forEach(({ pokemon }) => {
      console.log(pokemon.name);
    })
  } catch {
    console.log(`Invalid location area: ${locationAreaName}`);
  }
}
