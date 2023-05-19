fn get_input() -> &'static str {
    return "forward 5
down 5
forward 8
up 3
down 8
forward 2";
}

#[derive(Debug)]
struct Point {
    x: i32,
    y: i32,
}

fn parse_line(line: &'static str) -> Point {
    let (dir, amount) = line.split_once(" ").expect("must contain a whitespace");
    let amount = str::parse::<i32>(amount).expect("second arg must be an integer");

    if dir == "forward" {
        return Point { x: amount, y: 0 };
    } else if dir == "up" {
        return Point { y: -amount, x: 0 };
    }
    return Point { y: amount, x: 0 };
}

pub fn day_two() {
    let result =
        get_input()
            .lines()
            .map(|x| parse_line(x))
            .fold(Point { x: 0, y: 0 }, |mut acc, point| {
                acc.x += point.x;
                acc.y += point.y;
                return acc;
            });
    println!("{:?}", result)
}
