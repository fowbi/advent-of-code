use adventofcode_2023::run;

fn hand_type_bis(hand: &str) -> (&str, i32, &str, i32) {
    let diff_cards = vec!["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "K", "Q", "A"];
    let mut cards = vec![0; 13];

    hand.split("").for_each(|c| {
        match c {
            "2" => cards[0] += 1,
            "3" => cards[1] += 1,
            "4" => cards[2] += 1,
            "5" => cards[3] += 1,
            "6" => cards[4] += 1,
            "7" => cards[5] += 1,
            "8" => cards[6] += 1,
            "9" => cards[7] += 1,
            "T" => cards[8] += 1,
            "J" => cards[9] += 1,
            "Q" => cards[10] += 1,
            "K" => cards[11] += 1,
            "A" => cards[12] += 1,
            _ => (),
        }
    });

    if cards.iter().filter(|&&c| c == 5).count() == 1 {
        return (hand, 7, hand, 7);
    }

    if cards.iter().filter(|&&c| c == 4).count() == 1 {
        if cards[9] == 1 {
            let char = cards.iter().find(|&&c| c == 4).unwrap();
            println!("char: {} {:?} ", char, diff_cards.get(*char).unwrap());
        }
        //return (cards, 6);
    }

    if cards.iter().filter(|&&c| c == 3).count() == 1 {
        if cards.iter().filter(|&&c| c == 2).count() == 1 {
            //return (cards, 5);
        }
        //return (cards, 4);
    }

    if cards.iter().filter(|&&c| c == 2).count() == 2 {
        if cards[9] >= 1 {
            println!("cards: {:?}", cards);
            let char = cards.iter().position(|&c| c == 2).unwrap();
            println!("char: {}", char);
        }
        //return (cards, 3);
    }

    if cards.iter().filter(|&&c| c == 2).count() == 1 {
        //return (cards, 2);
    }

    (hand, 1, hand, 1)
}

fn hand_type(hand: &str) -> i32 {
    let mut cards = [0; 13];

    hand.split("").for_each(|c| {
        match c {
            "2" => cards[0] += 1,
            "3" => cards[1] += 1,
            "4" => cards[2] += 1,
            "5" => cards[3] += 1,
            "6" => cards[4] += 1,
            "7" => cards[5] += 1,
            "8" => cards[6] += 1,
            "9" => cards[7] += 1,
            "T" => cards[8] += 1,
            "J" => cards[9] += 1,
            "Q" => cards[10] += 1,
            "K" => cards[11] += 1,
            "A" => cards[12] += 1,
            _ => (),
        }
    });

    if cards.iter().filter(|&&c| c == 5).count() == 1 {
        return 7;
    }

    if cards.iter().filter(|&&c| c == 4).count() == 1 {
        return 6;
    }

    if cards.iter().filter(|&&c| c == 3).count() == 1 {
        if cards.iter().filter(|&&c| c == 2).count() == 1 {
            return 5;
        }
        return 4;
    }

    if cards.iter().filter(|&&c| c == 2).count() == 2 {
        return 3;
    }

    if cards.iter().filter(|&&c| c == 2).count() == 1 {
        return 2;
    }

    1
}

fn part1(input: &str) -> i32{
    let hands: Vec<(String, i32, i32)> = input.lines().map(|line| {
        let hand = line.split_whitespace().nth(0).unwrap();
        let number = line.split_whitespace().nth(1).unwrap().parse::<i32>().unwrap();
        let htype = hand_type(hand);
        let hand = format!("{}{}", htype, hand);

        (hand, number, htype)
    }).collect::<Vec<_>>();

    let mut shands = hands.iter().map(|x| {
        let y = x.0.split("").map(|r| match r {
            "K" => "B",
            "Q" => "C",
            "J" => "D",
            "T" => "E",
            "9" => "F",
            "8" => "G",
            "7" => "H",
            "6" => "I",
            "5" => "J",
            "4" => "K",
            "3" => "L",
            "2" => "M",
            _ => r,
        }).collect::<Vec<_>>().join("");
        (y, &x.0, &x.1, &x.2)
    }).collect::<Vec<_>>();
    shands.sort_by(|a, b| {
        if a.3 != b.3 {
            return (a.3).cmp(b.3);
        }
        b.0.to_lowercase().cmp(&a.0.to_lowercase())
    });
    shands.iter().enumerate().map(|(i, x)| x.2 * (i + 1) as i32).sum()
}

fn part2(input: &str) -> i64 {
    0
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 7 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 7 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483";

    #[rstest]
    #[case(INPUT_1, 6440)]
    fn star_07_part1(#[case] input: &str, #[case] expected_output: i32) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_1, 46)]
    fn star_07_part2(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part2(&input), expected_output);
    }
}
