use adventofcode_2023::run;
use std::collections::HashMap;
use regex::Regex;

fn extract_numbers_part_1(text: &str) -> Vec<char> {
    let mut numbers = Vec::new();

    for c in text.chars() {
        match c.is_ascii_digit() {
            true => numbers.push(c),
            false => continue,
        }
    }

    numbers
}

fn replace_number(text: &str) -> char {
    let textual_numbers = HashMap::from([
        ("one", "1"),
        ("two", "2"),
        ("three", "3"),
        ("four", "4"),
        ("five", "5"),
        ("six", "6"),
        ("seven", "7"),
        ("eight", "8"),
        ("nine", "9"),
    ]);

    for (key, value) in textual_numbers.into_iter() {
        let re = Regex::new(format!("^{}", key).as_str()).unwrap();
        if re.is_match(text) {
            return value.chars().next().unwrap();
        }
    }

    return text.chars().next().unwrap();
}

fn extract_numbers_part_2(text: &str) -> Vec<char> {
    let mut numbers = Vec::new();
    let mut text_copy = text.to_string();

    let mut textual_numbers = String::from("");
    while text_copy.len() > 0 {
        textual_numbers.push(replace_number(text_copy.as_str()));
        text_copy = text_copy[1..].to_string();
    }

    for c in textual_numbers.chars() {
        match c.is_ascii_digit() {
            true => numbers.push(c),
            false => continue,
        }
    }

    numbers
}

fn part1(input: &str) {
    let mut total = 0;
    for line in input.lines() {
        let numbers = extract_numbers_part_1(&line);
        total += format!("{}{}", numbers[0], numbers.last().unwrap()).parse::<i32>().unwrap();
    }
    println!("Day 1 Part 1: {}", total);
}

fn part2(input: &str) {
    let mut total = 0;
    for line in input.lines() {
        let numbers_part_2 = extract_numbers_part_2(&line);
        let c = format!("{}{}", numbers_part_2[0], numbers_part_2.last().unwrap()).parse::<i32>().unwrap();
        total += c;
    }
    println!("Day 1 Part 2: {}", total);
}

fn main() {
    run(part1);
    run(part2);
}
