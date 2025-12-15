use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Paragraph {
    body: String,
}

#[derive(Serialize, Deserialize)]
struct Article {
    article: String,
    author: String,
    paragraph: Vec<Paragraph>,
}

fn main() {
    let json = r#"
    {
        "article": "how to work with json in rust",
        "author": "birdman",
        "paragraph": [
            {
                "body": "starting sentence"
            },
            {
                "body": "middle sentence"
            },
            {
                "body": "ending sentence"
            }
        ]

    }
    "#;

    let parsed: Article = read_json_typed(json);
    println!("\n\n First paragraph is: {}", parsed.paragraph[0].body)
}

fn read_json_typed(raw_json: &str) -> Article {
    let parsed: Article = serde_json::from_str(raw_json).unwrap();
    parsed
}
