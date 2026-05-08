import { Cache } from "./pokecache.js"

const SECOND = 1000;

export class PokeAPI {
  private static readonly baseURL = "https://pokeapi.co/api/v2";
  cache: Cache;

  constructor() {
    this.cache = new Cache(30 * SECOND);
  }

  async fetchLocationAreas(pageURL?: string): Promise<ShallowLocationAreas> {
    const url = pageURL ?? `${PokeAPI.baseURL}/location-area`;
    const cacheData = this.cache.get<ShallowLocationAreas>(url);
    if (cacheData !== undefined) {
      return cacheData;
    }
    const data = await fetch(url);
    const shallowLocations: ShallowLocationAreas = await data.json();
    this.cache.add(url, shallowLocations);
    return shallowLocations;
  }

  async fetchLocationArea(locationName: string): Promise<LocationArea> {
    const url = `${PokeAPI.baseURL}/location-area/${locationName}`
    const cacheData = this.cache.get<LocationArea>(url);
    if (cacheData !== undefined) {
      return cacheData;
    }
    const data = await fetch(url);
    const locations: LocationArea = await data.json();
    this.cache.add(url, locations);
    return locations;
  }

  async fetchPokemon(name: string): Promise<PokemonData> {
    const url = `${PokeAPI.baseURL}/pokemon/${name}`
    const cacheData = this.cache.get<PokemonData>(url);
    if (cacheData !== undefined) {
      return cacheData;
    }
    const data = await fetch(url);
    const pokemonData: PokemonData = await data.json();
    this.cache.add(url, pokemonData);
    return pokemonData;
  }
}

export type ShallowLocationAreas = {
  count: number;
  next: string;
  previous: string;
  results: LocationAreaResult[];
}

type LocationAreaResult = {
  name: string;
  url: string;
}


export type LocationArea = {
  encounter_method_rates: EncounterMethodRate[];
  game_index: number;
  id: number;
  location: Location;
  name: string;
  names: Name[];
  pokemon_encounters: PokemonEncounter[];
}

type EncounterMethodRate = {
  encounter_method: EncounterMethod;
  version_details: VersionDetail[];
}

type EncounterMethod = {
  name: string;
  url: string;
}

type VersionDetail = {
  rate: number;
  version: Version;
}

type Version = {
  name: string;
  url: string;
}

type Location = {
  name: string;
  url: string;
}

type Name = {
  language: Language;
  name: string;
}

type Language = {
  name: string;
  url: string;
}

type PokemonEncounter = {
  pokemon: Pokemon;
  version_details: EncounterVersionDetail[];
}

type Pokemon = {
  name: string;
  url: string;
}

type EncounterVersionDetail = {
  encounter_details: EncounterDetail[];
  max_chance: number;
  version: Version;
}

type EncounterDetail = {
  chance: number;
  condition_values: any[];
  max_level: number;
  method: Method;
  min_level: number;
}

type Method = {
  name: string;
  url: string;
}

export type PokemonData = {
  id: number;
  name: string;
  base_experience: string;
}
