
type CacheEntry<T> = {
  createdAt: number;
  val: T;
}

export class Cache {
  #cache = new Map<string, CacheEntry<any>>();
  #reapIntervalId: NodeJS.Timeout | undefined = undefined;
  #interval: number;

  constructor(interval: number) {
    this.#interval = interval;
    this.#startReapLoop();
  }

  add<T>(key: string, val: T) {
    const entry: CacheEntry<T> = {
      createdAt: Date.now(),
      val: val,
    };
    this.#cache.set(key, entry);
  }

  get<T>(key: string): T | undefined {
    const entry = this.#cache.get(key);
    return entry?.val
  }

  stopReapLoop() {
    clearInterval(this.#reapIntervalId);
    this.#reapIntervalId = undefined;
  }


  #startReapLoop() {
    this.#reapIntervalId = setInterval(() => this.#reap(), this.#interval);
  }

  #reap() {
    this.#cache.forEach((value, key, map) => {
      if (Date.now() - value.createdAt > this.#interval) {
        // this is too old, needs to be reaped.
        map.delete(key);
      }
    })
  }
}

