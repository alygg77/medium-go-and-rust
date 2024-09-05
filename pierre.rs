use std::fs::File;
use std::io::{BufRead, BufReader};
use std::time::Instant;

fn main() {
    let start = Instant::now();
    let file = File::open("/Users/vladimir/Downloads/book-war-and-peace.txt").expect("Cannot open file");
    let reader = BufReader::new(file);




    // Initialize the counter for "Pierre"
    let mut count = 0;

    // Loop through each line in the file
    for line in reader.lines() {
        let line = line.expect("Error reading line");
        count += line.matches("Pierre").count();
    }

    // Calculate the elapsed time
    let duration = start.elapsed();

    // Output the result
    println!("Number of occurrences of 'Pierre': {}", count);
    println!("Time taken: {:?}", duration);
}
