use adventofcode_2023::run;
use std::collections::HashMap;
use regex::Regex;
pub use num::*;

fn map (input: &str) -> HashMap<String, (String, String)> {
    let mut elements = HashMap::new();
    input.lines().for_each(|line|{
        let re = Regex::new(r"(\w+) = \((\w+), (\w+)\)").unwrap();

        if !re.is_match(line) {
            return;
        }

        let cap = re.captures(line).unwrap();
        elements.insert(
            cap[1].to_string(),
            (cap[2].to_string(), cap[3].to_string()),
        );
    });

    elements
}

fn find_counts (starting_element: String, elements: HashMap<String, (String, String)>, directions: &str, ends_with: &str) -> i64 {
    let mut counter: i64 = 0;
    let mut element = starting_element.clone();
    for direction in directions.chars().cycle() {
        if direction == 'L' {
            element = elements[&element].0.to_string();
        } else {
            element = elements[&element].1.to_string();
        }

        counter += 1;
        if element.ends_with(ends_with) {
            break;
        }
    };

    counter
}

fn part1(input: &str) -> i64 {
    let directions = input.lines().next();
    let elements = map(input);
    let starting_element = "AAA".to_string();
    find_counts(starting_element, elements, directions.unwrap(),"ZZZ")
}

fn part2(input: &str) -> i64 {
    let directions = input.lines().next().unwrap();
    let elements = map(input);
    let starting_elements = elements.keys().filter(|&x| x.ends_with('A')).cloned().collect::<Vec<String>>();
    let mut counts = vec![];
    starting_elements.iter().for_each(|starting_element|{
        counts.push(find_counts(starting_element.to_string(), elements.clone(), directions, "Z"));
    });

    counts.into_iter().reduce(|a, b| a.lcm(&b)).unwrap()
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 8 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 8 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)";
    const INPUT_2: &str = "LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)";

    const INPUT_3: &str = "LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)";

    #[rstest]
    #[case(INPUT_1, 2)]
    #[case(INPUT_2, 6)]
    fn star_08_part1(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_3, 6)]
    fn star_08_part2(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part2(&input), expected_output);
    }
}
