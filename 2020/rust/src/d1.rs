fn get_input() -> &'static str {
    "1721
    979
    366
    299
    675
    1456"
}

fn parse_input(input: &str) -> Vec<u32> {
    let numbers: Vec<u32> = input.lines().map(|x| x.parse().unwrap()).collect();
    return numbers;
}

fn compare_numbers(numbers: Vec<u32>) -> u32 {
    let mut num1 = 0;
    let mut num2 = 0;
    for i in 0..numbers.len() {
        for j in i + 1..numbers.len() {
            if numbers[i] + numbers[j] == 2020 {
                num1 = numbers[i];
                num2 = numbers[j];
                break;
            }
        }
    }

    return num1 * num2;
}

// This is horrendous I know
fn compare_three_numbers(numbers: Vec<u32>) -> u32 {
    let mut num1 = 0;
    let mut num2 = 0;
    let mut num3 = 0;
    for i in 0..numbers.len() {
        for j in i + 1..numbers.len() {
            for k in j + 1..numbers.len() {
                if numbers[i] + numbers[j] + numbers[k] == 2020 {
                    num1 = numbers[i];
                    num2 = numbers[j];
                    num3 = numbers[k];
                    break;
                }
            }
        }
    }

    return num1 * num2 * num3;
}

pub fn day_one() {
    let input = get_input();
    println!("{}", compare_numbers(parse_input(input)));
    println!("{}", compare_three_numbers(parse_input(input)));
}
