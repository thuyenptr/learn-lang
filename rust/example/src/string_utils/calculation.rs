pub fn calculate_length(s: &String) -> usize {
    println!("string value {}", s);
    return s.len();
}

pub fn change_string(s: &mut String) -> usize {
    println!("string value {}", s);
    s.push_str(" ,hello");
    return s.len();
}