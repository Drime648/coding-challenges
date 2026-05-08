import { Cache } from "./pokecache.js";
import { test, expect } from "vitest";

test.concurrent.each([
  {
    key: "1",
    val: "test",
    interval: 500
  },
  {
    key: "2",
    val: "other",
    interval: 1000
  }
])("Test caching $interval ms", async ({ key, val, interval }) => {
  const cache = new Cache(interval);

  cache.add(key, val);
  const cached = cache.get(key);
  expect(cached).toBe(val);
  await new Promise((resolve) => setTimeout(resolve, interval * 2));
  const reaped = cache.get(key);
  expect(reaped).toBe(undefined);
  cache.stopReapLoop();
});
