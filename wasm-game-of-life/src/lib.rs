mod utils;

use std::slice::GetDisjointMutError;

use wasm_bindgen::prelude::*;

#[wasm_bindgen]
#[repr(u8)]
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum Cell {
    Dead = 0,
    Alive = 1,
}

#[wasm_bindgen]
pub struct Universe {
    width: u32,
    height: u32,
    cells: Vec<Cell>,
}

impl Universe {
    fn get_index(&self, row: u32, col: u32) -> usize {
        (row * self.width + col) as usize
    }

    fn num_neighbors(&self, row: u32, col: u32) -> u8 {
        let mut count: u8 = 0;
        for r in row - 1..row + 1 {
            for c in col - 1..col + 1 {
                if r == row && c == col {
                    continue;
                }
                count += self.cells[self.get_index(r, c)] as u8;
            }
        }
        return count;
    }
}

#[wasm_bindgen]
impl Universe {
    pub fn tick(&mut self) {
        let mut next = self.cells.clone();

        for row in 0..self.height {
            for col in 0..self.width {
                let idx = self.get_index(row, col);
                let num_neighbors = self.num_neighbors(row, col);
                match self.cells[idx] {
                    Cell::Alive => match num_neighbors {
                        0 | 1 => {
                            next[idx] = Cell::Dead;
                        }
                        2 | 3 => {
                            next[idx] = Cell::Alive;
                        }
                        _ => {
                            //above 3
                            next[idx] = Cell::Dead;
                        }
                    },
                    Cell::Dead => {
                        if num_neighbors == 3 {
                            next[idx] = Cell::Alive;
                        }
                    }
                }
            }
        }
        self.cells = next;
    }
}

// #[wasm_bindgen]
// extern "C" {
//     fn alert(s: &str);
// }

// #[wasm_bindgen]
// pub fn greet(name: &str) {
//     alert(format!("Hello, {}!", name).as_str());
// }
