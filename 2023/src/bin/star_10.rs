use adventofcode_2023::run;
pub use num::*;
use colored::Colorize;

fn map (input: &str) -> Vec<Vec<&str>> {
    let mut grid: Vec<Vec<&str>> = vec![];
    let len = input.lines().nth(0).unwrap().len() + 2;

    grid.push(vec!["."; len]);
    input.lines().for_each(|line| {
        let mut line = line.split("").filter(|s| s != &"").collect::<Vec<&str>>();
        line.insert(0, ".");
        line.push(".");
        grid.push(line);
    });
    grid.push(vec!["."; len]);

    grid
}

fn sp(input: &str) -> (usize, usize) {
    let starting_y = input.lines().position(|line| line.contains("S")).unwrap();
    (starting_y + 1, input.lines().nth(starting_y).unwrap().find("S").unwrap() + 1)
}

fn match_p (
    p: (usize, usize),
    pp: (usize, usize),
    directions: Vec<Vec<&str>>,
) -> ((usize, usize), Vec<(usize, usize)>, Vec<(usize, usize)>){
    //println!("p: {:?}", p);
    let u = directions[p.0-1][p.0];
    let l = directions[p.0][p.0-1];
    let r = directions[p.0][p.0+1];
    let d = directions[p.0+1][p.0];

    let mut points_down = vec![];
    let mut points_up = vec![];

    //println!("p: {:?} | char: {:?}", p, directions[p.0][p.1]);
    match directions[p.0][p.1] {
        "S" => {
            if u == "|" || u == "F" || u == "7" {
                //println!("Snext: {:?}", directions[p.0-1][p.1]);
                return ((p.0-1, p.1), vec![], vec![]);
            }
            if l == "-" || l == "L" || l == "F" {
                //println!("Snext: {:?}", directions[p.0][p.1-1]);
                return ((p.0, p.1-1), vec![], vec![]);
            }
            if r == "-" || r == "J" || r == "7" {
                //println!("Snext: {:?}", directions[p.0][p.1+1]);
                return ((p.0, p.1+1), vec![], vec![]);
            }
            //println!("Snext: {:?}", directions[p.0+1][p.1]);
            return ((p.0+1, p.1), vec![], vec![]);
        },
        "-" => {
            // for the . : always look up and down
            if directions[p.0-1][p.1] == "." {
                points_up.push((p.0+1, p.1));
            }
            if directions[p.0+1][p.1] == "." {
                points_down.push((p.0-1, p.1));
            }
            //println!("-: {} == {} && {} == {} >> ({} {}) >> {}", pp.0, p.0, pp.1, p.1, p.0, p.1+1, directions[p.0][p.1+1]);
            //println!("-: else ({} {}) >> {}", p.0, p.1-1, directions[p.0][p.1-1]);
            if pp.0 == p.0 && pp.1 == p.1-1 {
                return ((p.0, p.1+1), points_down, points_up);
            } else {
                return ((p.0, p.1-1), points_down, points_up);
            }
        },
        "|" => {
            // for the . : always left and right
            if directions[p.0][p.1-1] == "." {
                points_down.push((p.0, p.1-1));
            }
            if directions[p.0][p.1+1] == "." {
                points_up.push((p.0, p.1+1));
            }

            if pp.0 == p.0-1 && pp.1 == p.1 {
                return ((p.0+1, p.1), points_down, points_up);
            } else {
                return ((p.0-1, p.1), points_down, points_up);
            }
        },
        "7" => {
            // for the . : always right and up
            if directions[p.0][p.1+1] == "." {
                points_up.push((p.0, p.1+1));
            }
            if directions[p.0+1][p.1] == "." {
                points_up.push((p.0+1, p.1));
            }

            if pp.0 == p.0 && pp.1 == p.1-1 {
                return((p.0+1, p.1), points_down, points_up);
            } else {
                return((p.0, p.1-1), points_down, points_up);
            }
        },
        "F" => {
            // for the . : always left and up
            if pp.0 == p.0 && pp.1 == p.1+1 {
                return ((p.0+1, p.1), vec![], vec![]);
            } else {
                return ((p.0, p.1+1), vec![], vec![]);
            }
        },
        "J" => {
            if pp.0 == p.0-1 && pp.1 == p.1 {
                return ((p.0, p.1-1), vec![], vec![]);
            } else {
                return ((p.0-1, p.1), vec![], vec![]);
            }
        },
        "L" => {
            // for the . : always right and down
            if pp.0 == p.0-1 && pp.1 == p.1 {
                return ((p.0, p.1+1), vec![], vec![]);
            } else {
                return ((p.0-1, p.1), vec![], vec![]);
            }
        },
        _ => panic!("nope")
    }
}

fn part1(input: &str) -> i64 {
    let directions = map(input);
    let mut cp = sp(input);
    // -1,-1    |   -1,0    |   -1,3
    // 0,-1     |   0,0     |   0,1
    // 1,-1     |   1,0     |   1,1

    println!("starging point: {:?} {:?}", cp, directions[cp.0][cp.1]);
    let mut i = 0;
    let mut previous_p = cp.clone();
    let mut points: Vec<(usize, usize)> = vec![previous_p];
    let mut points_a: Vec<(usize, usize)> = vec![];
    let mut points_b: Vec<(usize, usize)> = vec![];
    loop {
        i+= 1;
        //if i == 100 {
            //break;
        //}
        let np = match_p(cp, previous_p, directions.clone());
        //println!("next point: {:?} {:?}", np, directions[np.0][np.1]);
        previous_p = cp;
        let r = np;
        cp = r.0;
        r.1.iter().for_each(|p| points_a.push(*p));
        r.2.iter().for_each(|p| points_b.push(*p));
        points.push(cp);

        if directions[cp.0][cp.1] == "S" {
            println!("{}", i);
            break;
        }
        //println!("cp: {:?} | pp {:?}", cp, previous_p);
    }

    directions.iter().enumerate().for_each(|(i,row)| {
        row.iter().enumerate().for_each(|(y, c)| {
            print!("{}", {
                if points.contains(&(i,y)) {
                    c.red()
                } else if points_a.contains(&(i,y)) {
                    "I".blue()
                } else if points_b.contains(&(i,y)) {
                    "0".green()
                } else {
                    c.normal().white()
                }
            });
        });
        println!("");
    });


    //println!("{:?}", points);

    i/2
}

fn part2(input: &str) -> i64 {
    let elements = map(input);
    0
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 10 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 10 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = ".....
.S-7.
.|.|.
.L-J.
.....";

    const INPUT_2: &str = "..F7.
.FJ|.
SJ.L7
|F--J
LJ...";

    #[rstest]
    #[case(INPUT_1, 4)]
    #[case(INPUT_2, 8)]
    fn star_10_part1(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part1(&input), expected_output);
    }

    //#[rstest]
    //#[case(INPUT_1, 0)]
    //fn star_10_part2(#[case] input: &str, #[case] expected_output: i64) {
        //assert_eq!(part2(&input), expected_output);
    //}
}
