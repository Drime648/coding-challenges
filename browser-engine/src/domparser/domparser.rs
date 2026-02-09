struct Node {
    children: Vec<NodeType>,
    node_type: NodeType,
}

enum NodeType {
    //node can either be raw text or an element
    Text(String),
    Element(ElementData),
}

struct ElementData {
    tag_name: String,
    attr_map: AttrMap,
}

type AttrMap = HashMap<String, String>;

impl Node {
    fn text(data: String) -> Node {
        Node {
            children: vec![],
            node_type: NodeType::Text(data),
        }
    }

    fn element(tag_name: String, attr_map: AttrMap, children: Vec<NodeType>) {
        Node {
            children,
            node_type: NodeType::ElementData(ElementData { tag_name, attr_map }),
        }
    }
}
