use adventofcode_2023::run;
use std::cmp::max;

fn get_game_number(game: &str) -> i32 {
    let digits = game
        .rsplit_once(|c: char| !c.is_ascii_digit())
        .map_or(game, |(_head, digits)| digits);
    digits.parse().unwrap()
}

fn get_cube_and_color(cube: &str) -> (i32, &str) {
    let mut parts = cube.trim().splitn(2, " ");
    let number_of_cubes = parts.next().unwrap().parse::<i32>().unwrap();
    let color = parts.next().unwrap().trim();
    (number_of_cubes, color)
}

fn parse_line(line: &str) -> (i32, Vec<(i32, i32, i32)>) {
    let mut parts = line.splitn(2, ":");
    let game = get_game_number(parts.next().unwrap().trim());
    let rounds = parts.next().unwrap().split(";");
    let mut played_rounds = Vec::new();
    for round in rounds {
        played_rounds.push(parse_round(round));
    }

    (game, played_rounds)
}

fn parse_round(round: &str) -> (i32, i32, i32) {
    let mut red = 0;
    let mut green = 0;
    let mut blue = 0;
    let parts = round.split(",");
    for part in parts {
        let (cube, color) = get_cube_and_color(part);
        match color {
            "red" => red += cube,
            "green" => green += cube,
            "blue" => blue += cube,
            _ => panic!("Unknown color"),
        }
    }
    (red, green, blue)
}

fn part1(input: &str) {
    let red_cubes = 12;
    let green_cubes = 13;
    let blue_cubes = 14;

    let mut total = 0;
    'outer: for line in input.lines() {
        let (game, rounds): (i32, Vec<(i32, i32, i32)>) = parse_line(line);
        for round in rounds {
            if round.0 > red_cubes || round.1 > green_cubes || round.2 > blue_cubes {
                println!("Game {} ({} {} {})({} {} {})", game, red_cubes, green_cubes, blue_cubes, round.0, round.1, round.2);
                continue 'outer;
            }
        }
        total += game;
    }

    println!("Day 2 Part 1: {}", total);
}

fn part2(input: &str) {
    let mut total = 0;
    for line in input.lines() {
        let (_, rounds): (i32, Vec<(i32, i32, i32)>) = parse_line(line);
        let mut minimum_red_cubes = 0;
        let mut minimum_green_cubes = 0;
        let mut minimum_blue_cubes = 0;
        for round in rounds {
            minimum_red_cubes = max(minimum_red_cubes, round.0);
            minimum_green_cubes = max(minimum_green_cubes, round.1);
            minimum_blue_cubes = max(minimum_blue_cubes, round.2);
        }
        total += minimum_red_cubes * minimum_green_cubes * minimum_blue_cubes;
    }

    println!("Day 2 Part 2: {}", total);
}

fn main() {
    run(part1);
    run(part2);
}
