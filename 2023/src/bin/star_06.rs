use adventofcode_2023::run;
use regex::Regex;

fn get_numbers(line: &str) -> Vec<i32> {
    Regex::new(r"\d+").unwrap().captures_iter(line).map(|cap| {
        cap[0].parse::<i32>().unwrap()
    }).collect()
}

fn get_number_as_str(line: &str) -> i64 {
    let regex = Regex::new(r"\d+").unwrap();
    let mut number = String::from("");
    for cap in regex.captures_iter(line) {
        number += &cap[0];
    }
    number.parse::<i64>().unwrap()
}

fn part1(input: &str) -> i32{
    let mut total = 1;

    let time = get_numbers(input.lines().nth(0).unwrap());
    let distance = get_numbers(input.lines().nth(1).unwrap());
    let races: Vec<(i32, i32)> = time.iter().zip(distance.iter()).map(|(t, d)| (*t, *d)).collect();

    races.iter().for_each(|(time, distance)| {
        let mut record_broken = 0;
        for i in 1..time-1 {
            let maximum_distance = i * (time - i);
            record_broken += if &maximum_distance > distance { 1 } else { 0 };
        }
        total *= record_broken;
    });

    total
}

fn part2(input: &str) -> i64{
    let mut total = 1;

    let races: Vec<(i64, i64)> = vec![(
        get_number_as_str(input.lines().nth(0).unwrap()),
        get_number_as_str(input.lines().nth(1).unwrap())
    )];

    println!("{:?}", races);

    races.iter().for_each(|(time, distance)| {
        let mut record_broken = 0;
        for i in 1..time-1 {
            let maximum_distance = i * (time - i);
            record_broken += if &maximum_distance > distance { 1 } else { 0 };
        }
        total *= record_broken;
    });

    total
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 6 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 6 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "Time:      7  15   30
Distance:  9  40  200";

    #[rstest]
    #[case(INPUT_1, 288)]
    fn star_06_part1(#[case] input: &str, #[case] expected_output: i32) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_1, 71503)]
    fn star_06_part2(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part2(&input), expected_output);
    }
}
