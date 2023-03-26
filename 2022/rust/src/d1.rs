fn get_input() -> &'static str {
    "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"
}

fn get_total_cals(elf: &str) -> u32 {
    let nums: Vec<&str> = elf.split("\n").collect();
    let mut total: u32 = 0;

    for num in nums {
        let val = match num.parse::<u32>() {
            Ok(n) => n,
            Err(_) => {
                println!("Failed to parse integer");
                0
            }
        };
        total += val;
    }

    total
}

//part 1

pub fn find_most_calories() -> u32 {
    let elves: Vec<&str> = get_input().split("\n\n").collect();
    let mut most_cals = 0;

    for elf in elves {
        let total_cals = get_total_cals(elf);

        if total_cals > most_cals {
            most_cals = total_cals
        }
    }

    most_cals
}

pub fn total_top_three() -> u32 {
    let mut elf_cal_totals: Vec<u32> = Vec::new();
    let elves: Vec<&str> = get_input().split("\n\n").collect();
    let mut most_cals = 0;

    for elf in elves {
        elf_cal_totals.push(get_total_cals(elf));
    }

    elf_cal_totals.sort_unstable();
    elf_cal_totals.reverse();

    most_cals += elf_cal_totals[0] + elf_cal_totals[1] + elf_cal_totals[2];

    most_cals
}
