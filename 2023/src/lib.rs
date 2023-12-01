use std::path::{Path, PathBuf};
use std::time::Instant;
use std::{env, fs};

pub fn run(f: impl Fn(&str)) {
    let input = get_input();
    println!("---");
    let start = Instant::now();
    f(&input);
    let duration = start.elapsed();
    println!("--- {duration:?}")
}

fn get_input() -> String {
    let bin = binary_name();
    let day = bin
        .strip_prefix("star_")
        .and_then(|b| b.parse::<u8>().ok())
        .unwrap();

    let path = make_path(&bin);
    match path.exists() {
        true => fs::read_to_string(path).map_err(anyhow::Error::from),
        false => panic!("No input file found for day {}", day),
    }
    .unwrap()
}

fn binary_name() -> String {
    env::args()
        .next()
        .as_ref()
        .map(Path::new)
        .and_then(Path::file_name)
        .and_then(std::ffi::OsStr::to_str)
        .map(String::from)
        .expect("No binary name found")

}

fn make_path(bin_name: &str) -> PathBuf {
    let mut path = PathBuf::from(env!("CARGO_MANIFEST_DIR"));

    path.push("input");
    path.push(bin_name);
    path.set_extension("txt");

    path
}
