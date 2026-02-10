mod domparser;

use std::collections::HashMap;

use domparser::types::Node;
use domparser::utils::print_node;

fn main() {
    let text_node = Node::text(String::from("hello world"));
    let main_node = Node::element(String::from("p"), HashMap::new(), vec![text_node]);
    print_node(&main_node, 0);
}
