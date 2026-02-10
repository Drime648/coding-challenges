use crate::domparser::types::{AttrMap, ElementData, Node, NodeType};

struct Parser {
    pos: usize,
    input: String,
}

impl Parser {
    fn next(&self) -> char {
        self.input[self.pos..].chars().next().unwrap()
    }

    fn next_is(&self, expected: char) -> bool {
        self.next() == expected
    }

    fn consume(&mut self) -> char {
        let c = self.next();
        self.pos += c.len_utf8();
        return c;
    }

    fn expect_next(&mut self, expected: char) {
        if self.next_is(expected) {
            self.consume();
        }
        panic!("expected char: {}", expected);
    }

    fn next_starts_with(&self, expected: &str) -> bool {
        self.input[self.pos..].starts_with(expected)
    }

    fn consume_while(&mut self, f: fn(&Self) -> bool) -> String {
        let mut output = String::new();
        while f(&self) {
            output.push(self.consume());
        }
        output
    }

    fn consume_until(&mut self, c: char) -> String {
        let mut output = String::new();
        while !self.next_is(c) {
            output.push(self.consume());
        }
        output
    }

    fn is_end(&self) -> bool {
        self.pos >= self.input.len()
    }

    fn consume_whitespace(&mut self) {
        self.consume_while(|s| s.next().is_whitespace());
    }

    fn consume_word(&mut self) -> String {
        self.consume_while(|s| !s.next().is_whitespace())
    }

    fn consume_string(&mut self) -> String {
        self.expect_next('"');
        let output = self.consume_until('"');
        self.expect_next('"');
        output
    }

    fn consume_attributes(&mut self) -> AttrMap {
        let mut attr_map = AttrMap::new();
        loop {
            self.consume_whitespace();
            if (self.next_is('>')) {
                break;
            }

            let attr_key = self.consume_while(|s| !s.next().is_whitespace() && !(s.next() == '='));
            self.consume_whitespace();
            let attr_value = self.consume_string();
            attr_map.insert(attr_key, attr_value);
        }
        attr_map
    }

    fn parse_tag(&mut self) -> ElementData {
        self.expect_next('<');
        let tag_name = self.consume_word();
        let attr_map = self.consume_attributes();
        ElementData { tag_name, attr_map }
    }

    // fn parse_string(&mut self) -> Node {
    //
    // }
}
