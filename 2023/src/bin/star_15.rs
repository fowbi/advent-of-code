use adventofcode_2023::run;
use linked_hash_map::LinkedHashMap;


fn hash_algo(word: &str) -> usize {
    let mut hash = 0;
    for c in word.chars() {
        hash = (hash + (c as usize)) * 17 % 256;
    }
    hash
}

fn part1(input: &str) -> usize {
    let words = input.split(",").collect::<Vec<&str>>();
    words.iter().map(|word| hash_algo(word.trim())).sum()
}

fn part2(input: &str) -> usize {
    let mut boxes = vec![LinkedHashMap::<&str, usize>::new(); 256];
    let words = input.split(",").collect::<Vec<&str>>();

    for word in words {
        let mut step = word.trim().split(|c| c == '=' || c == '-');
        let label = step.next().unwrap();
        let focal_length = step.next().unwrap();

        let box_id = hash_algo(label);

        if focal_length.is_empty() {
            boxes[box_id].remove(label);
        } else {
            *boxes[box_id].entry(label).or_insert(0) = focal_length.parse::<usize>().unwrap();
        }
    }

    let mut sum = 0;
    for (box_id, box_content) in boxes.iter().enumerate() {
        for lens_id in 0..box_content.len() {
            let focal_length = box_content.keys().nth(lens_id).unwrap();
            // lens_box_number + slot_number + focal_length
            sum += (box_id + 1) * (lens_id + 1) * boxes[box_id][focal_length] as usize;
        }
    }

    sum
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 15 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 15 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7";

    #[rstest]
    #[case(INPUT_1, 1320)]
    fn star_15_part1(#[case] input: &str, #[case] expected_output: usize) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_1, 145)]
    fn star_15_part2(#[case] input: &str, #[case] expected_output: usize) {
        assert_eq!(part2(&input), expected_output);
    }
}
