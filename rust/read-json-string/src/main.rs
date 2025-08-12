use std::error::Error;
use std::fs::read_to_string;
use serde_json::{Value, from_str};

fn main() -> Result<(), Box<dyn Error>> {
    let instance_string = read_to_string("data/instance.json")?;
    let json: Value = from_str(&instance_string)?;

    if let Some(object) = json.as_object() {
        for (key, value) in object.iter() {
            println!("{key}: {value}");
        }
    }

    Ok(())
}