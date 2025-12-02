use dotenv::from_filename;
use std::env;

fn main() -> anyhow::Result<()> {
    println!("Hello, world!");
    from_filename(".env").ok();

    let wallet_descriptor = env::var("WALLET_DESCRIPTOR")?;
    // println!("Descriptor: {}", wallet_descriptor);

    dbg!(wallet_descriptor);
    Ok(())
}
