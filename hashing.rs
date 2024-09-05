use ed25519_dalek::{Keypair, Signer};
use rand::rngs::OsRng;
use std::time::Instant;

fn main() {
    // Create a new keypair using a cryptographic random number generator
    let mut csprng = OsRng;
    let keypair = Keypair::generate(&mut csprng);

    // Message to sign
    let message = b"Hello, World!";

    // Measure time to sign
    let total_iterations = 100000;
    let mut durations = vec![];

    for _ in 0..total_iterations {
        let start = Instant::now();
        let _signature = keypair.sign(message);
        let duration = start.elapsed();
        durations.push(duration);
    }

    // Calculate total time and average time
    let total_time: std::time::Duration = durations.iter().sum();
    let avg_time = total_time / total_iterations as u32;

    println!("Average hashing/signing time: {:?}", avg_time);
    println!("Total time for {} iterations: {:?}", total_iterations, total_time);
}
