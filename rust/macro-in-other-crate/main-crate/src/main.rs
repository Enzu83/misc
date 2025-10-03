use macro_crate::exec_structure_func;


fn main() {
    let sample1 = Sample{ name: "sample1".to_string() };
    let sample2 = Sample{ name: "sample2".to_string() };

    exec_structure_func!(sample2);
    // Method 3: Pass struct type for type-level operations
    exec_structure_func!(Sample);
}

pub struct Sample {
    pub name: String
}

pub trait Exec {
    fn func(&self);
}

impl Exec for Sample {
    fn func(&self) {
        println!("Hi from {}", self.name);
    }
}
