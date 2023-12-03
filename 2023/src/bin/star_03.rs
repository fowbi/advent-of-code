use adventofcode_2023::run;
use regex::Regex;

fn contains_only_digits(input: &str) -> bool {
    input.chars().all(|c| c.is_digit(10) || c == '.')
}

fn part1(input: &str) {
    let mut total = 0;
    for (i, line) in input.lines().enumerate() {
        let re = Regex::new(r"\d+").unwrap();
        re.captures_iter(line).for_each(|cap| {
            let start = cap.get(0).map(|t| t.start()).unwrap();
            let start = if start > 0 { start - 1 } else { start };
            let end = cap.get(0).map(|t| t.end()).unwrap();
            let end = if end < line.len() - 1 { end+1 } else { end };
            let value = cap.get(0).map(|t| t.as_str()).unwrap().parse::<i32>().unwrap();

            let mut check = "".to_owned();
            check.push_str(&line[start..end]);
            if i > 0 {
                check.push_str(&input.lines().nth(i-1).unwrap()[start..end]);
            }
            if i < input.lines().count() - 1 {
                check.push_str(&input.lines().nth(i+1).unwrap()[start..end]);
            }

            if !contains_only_digits(&check) {
                total += value;
            }
        });
    };

    println!("Day 2 Part 1: {}", total);
}

fn clean_up_line(input: &str) -> String {
    let mut output = String::new();

    for c in input.chars() {
        if c.is_digit(10) || c == '*' || c == '.' {
            output.push(c);
        } else {
            output.push('.');
        }
    }

    output
}

fn part2(input: &str) {
    let mut total = 0;

    let lines: Vec<String> = input.lines().map(|line| String::from(clean_up_line(line))).collect();
    lines.iter().enumerate().for_each(|(i, line)| {
        let previous_line = if i == 0 { line.replace(|_line| true, ".") } else { lines.iter().nth(i - 1).unwrap().to_string() };
        let next_line = if i == lines.iter().count() - 1 { line.replace(|_line| true, ".") } else { lines.iter().nth(i + 1).unwrap().to_string() };

        let mut gears = vec![];
        let re = Regex::new(r"\*").unwrap();
        re.captures_iter(line).for_each(|cap| {
            let start = cap.get(0).map(|t| t.start()).unwrap();
            let start = if start > 0 { start - 1 } else { start };
            let end = cap.get(0).map(|t| t.end()).unwrap();
            let end = if end < line.len() - 1 { end } else { end - 1 };
            gears.push((start, end, 0, 1));
        });


        for gear in gears.iter_mut() {
            let re = Regex::new(r"\d+").unwrap();
            re.captures_iter(previous_line.as_str()).for_each(|cap| {
                let start = cap.get(0).map(|t| t.start()).unwrap();
                let end = cap.get(0).map(|t| t.end()).unwrap() - 1;
                let value = cap.get(0).map(|t| t.as_str()).unwrap().parse::<i32>().unwrap();
                

                if (start <= gear.0 && end >= gear.0) || (start <= gear.1 && end >= gear.1) {
                    gear.2 += 1;
                    gear.3 *= value;
                }
            });
            re.captures_iter(line.as_str()).for_each(|cap| {
                let start = cap.get(0).map(|t| t.start()).unwrap();
                let end = cap.get(0).map(|t| t.end()).unwrap() - 1;
                let value = cap.get(0).map(|t| t.as_str()).unwrap().parse::<i32>().unwrap();

                if (start <= gear.0 && end >= gear.0) || (start <= gear.1 && end >= gear.1) {
                    gear.2 += 1;
                    gear.3 *= value;
                }
            });
            re.captures_iter(next_line.as_str()).for_each(|cap| {
                let start = cap.get(0).map(|t| t.start()).unwrap();
                let end = cap.get(0).map(|t| t.end()).unwrap() - 1;
                let value = cap.get(0).map(|t| t.as_str()).unwrap().parse::<i32>().unwrap();
                
                if (start >= gear.0 && end <= gear.0) || (start >= gear.1 && end <= gear.1) || (gear.1 >= start && gear.0 <= end) {
                    gear.2 += 1;
                    gear.3 *= value;
                }
            });

            if gear.2 == 2 {
                total += gear.3;
            }
        }
    });

    println!("Day 2 Part 2: {}", total);
}

fn main() {
    run(part1);
    run(part2);
}
