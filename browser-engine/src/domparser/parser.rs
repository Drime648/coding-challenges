use std::collections::HashMap;

pub struct Node {
    pub children: Vec<Node>,
    pub node_type: NodeType,
}

pub enum NodeType {
    //node can either be raw text or an element
    Text(String),
    Element(ElementData),
}

pub struct ElementData {
    pub tag_name: String,
    pub attr_map: AttrMap,
}

type AttrMap = HashMap<String, String>;

impl Node {
    fn text(data: String) -> Node {
        Node {
            children: vec![],
            node_type: NodeType::Text(data),
        }
    }

    fn element(tag_name: String, attr_map: AttrMap, children: Vec<NodeType>) -> Node {
        Node {
            children,
            node_type: NodeType::Element(ElementData { tag_name, attr_map }),
        }
    }
}
