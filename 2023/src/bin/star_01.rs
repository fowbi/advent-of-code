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

fn part1(input: &str) ->i32 {
    let mut total = 0;
    for line in input.lines() {
        let numbers = extract_numbers_part_1(&line);
        total += format!("{}{}", numbers[0], numbers.last().unwrap()).parse::<i32>().unwrap();
    }
    total
}

fn part2(input: &str) -> i32 {
    let mut total = 0;
    for line in input.lines() {
        let numbers_part_2 = extract_numbers_part_2(&line);
        let c = format!("{}{}", numbers_part_2[0], numbers_part_2.last().unwrap()).parse::<i32>().unwrap();
        total += c;
    }
    total
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 1 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 1 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";
    const INPUT_2: &str = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";


    #[rstest]
    #[case(INPUT_1, 142)]
    fn star_01_part1(#[case] input: &str, #[case] expected_output: i32) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_1, 142)]
    #[case(INPUT_2, 281)]
    fn start_01_part2(#[case] input: &str, #[case] expected_output: i32) {
        assert_eq!(part2(&input), expected_output);
    }
}
