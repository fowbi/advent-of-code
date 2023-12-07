use adventofcode_2023::run;
use std::ops::Range;
use std::thread;
use std::clone::Clone;

#[derive(Clone)]
struct Almenac {
    seeds: Vec<i64>,
    seed_to_soil_map: Vec<Vec<i64>>,
    soil_to_fertilizer_map: Vec<Vec<i64>>,
    fertilizer_to_water_map: Vec<Vec<i64>>,
    water_to_light_map: Vec<Vec<i64>>,
    light_to_temperature_map: Vec<Vec<i64>>,
    temperature_to_humidity_map: Vec<Vec<i64>>,
    humidity_to_location_map: Vec<Vec<i64>>,
}

fn parse_almenac(input: &str) -> Almenac {
    let mut lines = input.lines();

    let mut seed_to_soil_map: Vec<Vec<i64>> = Vec::new();
    let mut soil_to_fertilizer_map: Vec<Vec<i64>> = Vec::new();
    let mut fertilizer_to_water_map: Vec<Vec<i64>> = Vec::new();
    let mut water_to_light_map: Vec<Vec<i64>> = Vec::new();
    let mut light_to_temperature_map: Vec<Vec<i64>> = Vec::new();
    let mut temperature_to_humidity_map: Vec<Vec<i64>> = Vec::new();
    let mut humidity_to_location_map: Vec<Vec<i64>> = Vec::new();

    let seeds = lines.next().unwrap().split(": ").nth(1).unwrap().split(" ").map(|s| s.parse::<i64>().unwrap()).collect::<Vec<i64>>();

    let mut mapper = String::new();
    for line in lines {
        match line {
            "seed-to-soil map:" | "soil-to-fertilizer map:" | "fertilizer-to-water map:" | "water-to-light map:" | "light-to-temperature map:" | "temperature-to-humidity map:" | "humidity-to-location map:" => {
                mapper = line.to_string();
                continue;
            },
            s => {
                if s == "" {
                    continue;
                }
                let mut v: Vec<i64> = Vec::new();
                for value in s.trim().split(' ') {
                    v.push(value.parse().unwrap());
                }

                match mapper.as_str() {
                    "seed-to-soil map:" => seed_to_soil_map.push(v),
                    "soil-to-fertilizer map:" => soil_to_fertilizer_map.push(v),
                    "fertilizer-to-water map:" => fertilizer_to_water_map.push(v),
                    "water-to-light map:" => water_to_light_map.push(v),
                    "light-to-temperature map:" => light_to_temperature_map.push(v),
                    "temperature-to-humidity map:" => temperature_to_humidity_map.push(v),
                    "humidity-to-location map:" => humidity_to_location_map.push(v),
                    _ => unreachable!(),
                }
            }
        }
    }

    Almenac {
        seeds,
        seed_to_soil_map,
        soil_to_fertilizer_map,
        fertilizer_to_water_map,
        water_to_light_map,
        light_to_temperature_map,
        temperature_to_humidity_map,
        humidity_to_location_map,
    }
}

fn get_mapped(s: i64, mapping: Vec<Vec<i64>>) -> i64 {
    let mut mapped = s;
    mapping.iter().for_each(|m| {
        if (&s >= &m[1]) && (&s <= &(m[1] + m[2] - 1)) {
            mapped = s + m[0] - m[1];
        }
    });

    mapped
}

fn find_location(al: &Almenac, s: i64) -> i64 {
    let mut range = s;

    range = get_mapped(range, al.seed_to_soil_map.clone());
    range = get_mapped(range, al.soil_to_fertilizer_map.clone());
    range = get_mapped(range, al.fertilizer_to_water_map.clone());
    range = get_mapped(range, al.water_to_light_map.clone());
    range = get_mapped(range, al.light_to_temperature_map.clone());
    range = get_mapped(range, al.temperature_to_humidity_map.clone());
    range = get_mapped(range, al.humidity_to_location_map.clone());

    range
}

fn part1(input: &str) -> i64{
    let mut lowest = 0;
    let al = parse_almenac(input);

    al.seeds.iter().for_each(|&s| {
        let location = find_location(&al, s);
        if lowest == 0 || location < lowest {
            lowest = location;
        }
    });

    lowest
}

fn find_lowest_location(al: &Almenac, range: Range<i64>) -> i64 {
    range
        .into_iter()
        .map(|seed| find_location(&al, seed))
        .min()
        .unwrap()
}

fn part2(input: &str) -> i64 {
    let almenac = parse_almenac(input);

    let mut threads = vec![];

    let mut seed_ranges = vec![];
    let mut range_start = 0;
    for (i, &seed) in almenac.seeds.iter().enumerate() {
        if i % 2 == 0 {
            range_start = seed;
        } else {
            seed_ranges.push(range_start..(range_start + seed));
        }
    }

    for (_i, seed_range) in seed_ranges.into_iter().enumerate() {
        let almenac = almenac.clone();
        let thread = thread::spawn(move || {
            find_lowest_location(&almenac, seed_range.clone())
        });
        threads.push(thread);
    }

    let mut lowest_location = i64::MAX;
    for thread in threads {
        let result = thread.join().unwrap();
        if result < lowest_location {
            lowest_location = result;
        }
    }

    lowest_location
}

fn main() {
    fn run_part1(input: &str) {
        println!("Day 5 Part 1: {}", part1(input));
    }
    run(run_part1);

    fn run_part2(input: &str) {
        println!("Day 5 Part 2: {}", part2(input));
    }
    run(run_part2);
}

#[cfg(test)]
mod tests {
    use rstest::rstest;
    use super::{part1, part2};

    const INPUT_1: &str = "seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4";

    #[rstest]
    #[case(INPUT_1, 35)]
    fn star_05_part1(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part1(&input), expected_output);
    }

    #[rstest]
    #[case(INPUT_1, 46)]
    fn star_05_part2(#[case] input: &str, #[case] expected_output: i64) {
        assert_eq!(part2(&input), expected_output);
    }
}
