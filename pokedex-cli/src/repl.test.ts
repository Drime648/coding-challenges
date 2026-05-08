import { cleanInput } from "./repl.js";
import { describe, expect, test } from "vitest";

describe.each([
  {
    input: "  hello   world  ",
    expected: ["hello", "world"]
  },
  {
    input: "  ",
    expected: []
  }
])("cleanInput(%s)", ({ input, expected }) => {
  test(`Expected: ${expected}`, () => {
    let actual = cleanInput(input);
    expect(actual).toHaveLength(expected.length);
    for (const i in expected) {
      expect(actual[i]).toBe(expected[i])
    }
  });
});
