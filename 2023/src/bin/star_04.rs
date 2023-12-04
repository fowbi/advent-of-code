use adventofcode_2023::run;

fn get_numbers(line: &str) -> (Vec<&str>, Vec<&str>) {
    let (_, line) = line.split_at(line.find(':').unwrap() + 2);
    let mut iter = line.splitn(2, '|');
    (iter.next().unwrap().split_whitespace().collect(), iter.next().unwrap().split_whitespace().collect())
}

fn part1(input: &str) {
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

    println!("Day 2 Part 1: {}", total);
}

fn part2(input: &str) {
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

    println!("Day 2 Part 2: {}", num_of_cards.iter().fold(0, |acc, x| acc + x));
}

fn main() {
    run(part1);
    run(part2);
}
