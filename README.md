This is the code I've used for my Medium article

Rust info:
[package]
name = "Test"
version = "0.1.0"
edition = "2021"

[dependencies]
reqwest = { version = "0.11", features = ["json"] }
tokio = { version = "1", features = ["full"] }
serde_json = "1.0.127"
actix-web = "4.0"
ed25519-dalek = "1.0.1"
rand = { version = "0.7.0", features = ["std"] }
rand_core = "0.6"
