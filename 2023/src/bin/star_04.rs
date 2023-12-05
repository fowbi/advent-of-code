use adventofcode_2023::run;

fn get_numbers(line: &str) -> (Vec<&str>, Vec<&str>) {
    let (_, line) = line.split_at(line.find(':').unwrap() + 2);
    let mut iter = line.splitn(2, '|');
    (iter.next().unwrap().split_whitespace().collect(), iter.next().unwrap().split_whitespace().collect())
}

fn part1(input: &str) ->i32 {
    let mut total = 0;

    for line in input.lines() {
        let (winning_numbers, my_numbers) = get_numbers(line);
        let mut round_win = 0;
        winning_numbers.iter().for_each(|n| {
            if my_numbers.contains(n) {
                round_win = if round_win == 0 { 1 } else { round_win * 2 };
            }
        });

        total += round_win;
    };

    total
}

fn part2(input: &str) -> i32 {
    let mut num_of_cards = vec![1; input.lines().count()];
    for (card, line) in input.lines().enumerate() {
        let (winning_numbers, my_numbers) = get_numbers(line);
        let mut wins = 0;
        winning_numbers.iter().enumerate().for_each(|(_, n)| {
            if my_numbers.contains(n) {
                wins+= 1;
            }
        });
        for i in 0..wins {
            num_of_cards[card + 1 + i] += 1 * num_of_cards[card];
        }
    }

    num_of_cards.iter().fold(0, |acc, x| acc + x)
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 4 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 4 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";

    #[rstest]
    #[case(INPUT_1, 13)]
    fn star_02_part1(#[case] input: &str, #[case] expected_output: i32) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_1, 30)]
    fn start_02_part2(#[case] input: &str, #[case] expected_output: i32) {
        assert_eq!(part2(&input), expected_output);
    }
}
