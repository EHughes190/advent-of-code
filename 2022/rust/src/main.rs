pub mod d1;

fn main() {
    let most_cals = d1::find_most_calories();
    println!("Elf carrying the most calories: {}", most_cals);

    let total_top_three_cals = d1::total_top_three();
    println!("Total of top three elves: {}", total_top_three_cals);
}
