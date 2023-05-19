fn get_input() -> &'static str {
    return "1-3 a: abcde
    1-3 b: cdefg
    2-9 c: ccccccccc";
}

fn parse_lines(input: &str) -> u32 {
    let lines: Vec<(Vec<u32>, &str, &str)> = input.lines().map(|line| parse_line(line)).collect();
    let mut valid_passwords = 0;

    for line in lines {
        valid_passwords += verify_password(&line);
    }

    return valid_passwords;
}

fn parse_line(line: &str) -> (Vec<u32>, &str, &str) {
    let parts: Vec<&str> = line.split(" ").collect();
    let numbers: Vec<u32> = parts[0].split("-").map(|x| x.parse().unwrap()).collect();
    let key_letter = parts[1].split(":").next().unwrap();
    let result: (Vec<u32>, &str, &str) = (numbers, key_letter, parts[2]);

    return result;
}

fn verify_password(password_data: &(Vec<u32>, &str, &str)) -> u32 {
    let (bounds, key_letter, password) = password_data;
    let upper_bound = bounds[1];
    let lower_bound = bounds[0];
    let mut total = 0;
    let key = key_letter.chars().next().expect("string is empty");

    for char in password.chars() {
        if char == key {
            total += 1;
        }
    }

    if total <= upper_bound && total >= lower_bound {
        return 1;
    }

    return 0;
}

pub fn day_two() {
    println!("{}", parse_lines(get_input()));
}
