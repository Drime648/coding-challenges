use crate::domparser::parser::{Node, NodeType};

pub fn print_node(input: &Node, level: i32) {
    print_offset(level);
    print_node_type(&input.node_type);
    for node in &input.children {
        print_node(&node, level + 1);
    }
}

fn print_offset(level: i32) {
    for _ in 0..level {
        print!("\t");
    }
}

fn print_node_type(input: &NodeType) {
    match input {
        NodeType::Text(text) => {
            print!("{}", text);
        }
        NodeType::Element(elementData) => {
            print!("<{}", elementData.tag_name);
            for (key, value) in &elementData.attr_map {
                print!(" {} = {}", key, value);
            }
            print!(">");
        }
    }
    println!();
}
