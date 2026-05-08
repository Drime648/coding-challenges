import { State } from "./state.js";
import { ShallowLocationAreas } from "./pokeapi.js";

export async function commandMap(state: State, ..._args: string[]) {
  let locationAreas: ShallowLocationAreas;
  if (state.nextLocationsURL == null) {
    locationAreas = await state.pokeApi.fetchLocationAreas()
  } else {
    locationAreas = await state.pokeApi.fetchLocationAreas(state.nextLocationsURL)
  }
  locationAreas.results.forEach(({ name }) => {
    console.log(name);
  })
  state.nextLocationsURL = locationAreas.next;
  state.prevLocationsURL = locationAreas.previous;
}

export async function commandMapBack(state: State, ..._args: string[]) {
  let locationAreas: ShallowLocationAreas;
  if (state.prevLocationsURL == null) {
    locationAreas = await state.pokeApi.fetchLocationAreas()
  } else {
    locationAreas = await state.pokeApi.fetchLocationAreas(state.prevLocationsURL)
  }
  locationAreas.results.forEach(({ name }) => {
    console.log(name);
  })
  state.nextLocationsURL = locationAreas.next;
  state.prevLocationsURL = locationAreas.previous;
}
