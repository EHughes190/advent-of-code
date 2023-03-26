fn get_input() -> &'static str {
    "A Y
B X
C Z"
}

fn parse_rounds(l: &str) -> i32 {
    let parts: Vec<&str> = l.split(" ").collect();
    let po = parts[0];
    let pt = parts[1];

    match po {
        "A" => match pt {
            "X" => 1 + 3,
            "Y" => 2 + 6,
            "Z" => 3,
            _ => -1,
        },
        "B" => match pt {
            "X" => 1,
            "Y" => 2 + 3,
            "Z" => 3 + 6,
            _ => -1,
        },
        "C" => match pt {
            "X" => 1 + 6,
            "Y" => 2,
            "Z" => 3 + 3,
            _ => -1,
        },
        _ => -1,
    }
}

fn parse_rounds_part_two(l: &str) -> i32 {
    let parts: Vec<&str> = l.split(" ").collect();
    let po = parts[0];
    let pt = parts[1];

    match po {
        "A" => match pt {
            "X" => 3,
            "Y" => 1 + 3,
            "Z" => 2 + 6,
            _ => -1,
        },
        "B" => match pt {
            "X" => 1,
            "Y" => 2 + 3,
            "Z" => 3 + 6,
            _ => -1,
        },
        "C" => match pt {
            "X" => 2,
            "Y" => 3 + 6,
            "Z" => 1 + 6,
            _ => -1,
        },
        _ => -1,
    }
}

pub fn calculate_score() {
    let rounds: Vec<&str> = get_input().split("\n").collect();
    let mut total = 0;

    for r in rounds {
        total += parse_rounds(r);
    }

    println!("score for rock paper scissors: {}", total);
}

pub fn calculate_score_part_two() {
    let rounds: Vec<&str> = get_input().split("\n").collect();
    let mut total = 0;

    for r in rounds {
        total += parse_rounds_part_two(r);
    }

    println!("score for rock paper scissors part two: {}", total);
}
