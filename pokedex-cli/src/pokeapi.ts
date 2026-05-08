import { Cache } from "./pokecache.js"

const SECOND = 1000;

export class PokeAPI {
  private static readonly baseURL = "https://pokeapi.co/api/v2/location-area";
  cache: Cache;

  constructor() {
    this.cache = new Cache(30 * SECOND);
  }

  async fetchLocationAreas(pageURL?: string): Promise<ShallowLocationAreas> {
    const url = pageURL ?? PokeAPI.baseURL;
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
    const url = `${PokeAPI.baseURL}/${locationName}`
    const cacheData = this.cache.get<LocationArea>(url);
    if (cacheData !== undefined) {
      return cacheData;
    }
    const data = await fetch(url);
    const locations: LocationArea = await data.json();
    this.cache.add(url, locations);
    return locations;
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

