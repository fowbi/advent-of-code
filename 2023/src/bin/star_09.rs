use adventofcode_2023::run;

fn extract_numbers(input: &str) -> Vec<Vec<i64>> {
    input.lines().map(|line| {
        return line
            .split_whitespace()
            .map(|s| s.parse::<i64>().unwrap())
            .collect()
    }).collect::<Vec<_>>()
}

fn get_sequences(line: Vec<i64>, mut sequences: Vec<Vec<i64>>) -> Vec<Vec<i64>> {
    let mut line = line.clone();
    let mut last = line.pop().unwrap();
    let mut new_line: Vec<i64> = Vec::new();

    while !line.is_empty() {
        let next = line.pop().unwrap();
        let diff = last - next;
        last = next;
        new_line.push(diff);
    }
    new_line.reverse();
    sequences.push(new_line.clone());

    if new_line[new_line.len()-1] != 0 {
        return get_sequences(new_line, sequences);
    }

    sequences
}

fn part1(input: &str) -> i64 {
    extract_numbers(input)
        .iter()
        .map(|line| {
            get_sequences(line.clone(), vec![line.clone()])
                .iter()
                .rev()
                .fold(0, |a, b| a + b.last().unwrap())
        })
        .sum()
}

fn part2(input: &str) -> i64 {
    extract_numbers(input)
        .iter()
        .map(|line| {
            get_sequences(line.clone(), vec![line.clone()])
                .iter()
                .rev()
                .fold(0, |a, b| b[0] - a)
        })
        .sum()
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 9 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 9 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45";

    #[rstest]
    #[case(INPUT_1, 114)]
    fn star_09_part1(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_1, 2)]
    fn star_09_part2(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part2(&input), expected_output);
    }
}
