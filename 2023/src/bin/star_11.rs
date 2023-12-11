use adventofcode_2023::run;
use adventofcode_2023::Grd;
pub use num::*;

fn map(input: &str) -> Grd<char> {
    let mut lines = input.lines().clone();
    let mut grid = Grd::new(lines.nth(0).unwrap().len(), lines.count()+1);

    for (y, line) in input.lines().enumerate() {
        for (x, c) in line.chars().enumerate() {
            grid.set(x, y, c);
        }
    }

    grid
}

fn get_galaxies(grid: Grd<char>, offset: i64) -> Vec<(usize, usize)> {
    let mut empty_x_lines = vec![];
    let mut empty_y_lines = vec![];
    for x in 0..grid.width {
        let mut empty_line = false;
        for y in 0..grid.height {
            if grid.at(x, y).unwrap() == &'#' {
                empty_line = true;
            }
        }
        if !empty_line {
            empty_x_lines.push(x);
        }
    }

    for y in 0..grid.height {
        let mut empty_line = false;
        for x in 0..grid.width {
            if grid.at(x, y).unwrap() == &'#' {
                empty_line = true;
            }
        }
        if !empty_line {
            empty_y_lines.push(y);
        }
    }

    let mut galaxies = Vec::new();
    for y in 0..grid.height {
        for x in 0..grid.width {
            if grid.at(x, y).unwrap() == &'#' {

                let mut x_offset = offset;
                for ex in 0..empty_x_lines.len() {
                    if empty_x_lines[ex] < x {
                        x_offset += offset - 1;
                    }
                }

                let mut y_offset = offset;
                for ey in 0..empty_y_lines.len() {
                    if empty_y_lines[ey] < y {
                        y_offset += offset - 1;
                    }
                }

                galaxies.push((x + x_offset as usize, y + y_offset as usize));
            }
        }
    }

    galaxies
}

fn calc(galaxies: Vec<(usize, usize)>) -> i64 {
    let mut total_length = 0;
    for galaxy_y in 0..galaxies.len() {
        for galaxy_x in 0..galaxy_y {
            // taxicab: |p1 - q1| + |p2 - q2|
            let dx = galaxies[galaxy_y].0 as i64 - galaxies[galaxy_x].0 as i64;
            let dy = galaxies[galaxy_y].1 as i64 - galaxies[galaxy_x].1 as i64;
            total_length += abs(dx) + abs(dy)
        }
    }

    total_length
}

fn part1(input: &str) -> i64 {
    let grid = map(input);
    let galaxies = get_galaxies(grid.clone(), 2);
    calc(galaxies)
}

fn part2(input: &str) -> i64 {
    let grid = map(input);
    let galaxies = get_galaxies(grid.clone(), 1000000);
    calc(galaxies)
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 11 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 11 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....";

    #[rstest]
    #[case(INPUT_1, 374)]
    fn star_11_part1(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part1(&input), expected_output);
    }

    //#[rstest]
    //#[case(INPUT_1, 0)]
    //fn star_11_part2(#[case] input: &str, #[case] expected_output: i64) {
        //assert_eq!(part2(&input), expected_output);
    //}
}
