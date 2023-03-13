function getInput() {
  return `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`;
}

function getTotalCals(elf: string): number {
  const nums = elf.split("\n");
  let total = 0;

  for (let i = 0; i < nums.length; i++) {
    total += parseInt(nums[i]);
  }
  return total;
}

//Part 1
/* function main() { */
/*   const elves = getInput().split("\n\n"); */
/*   let mostCals = 0; */
/**/
/*   for (let i = 0; i < elves.length; i++) { */
/*     let totalCals = getTotalCals(elves[i]); */
/*     if (totalCals > mostCals) { */
/*       mostCals = totalCals; */
/*     } */
/*   } */
/*   console.log(mostCals); */
/* } */

function main() {
  const elves = getInput().split("\n\n");
  const elfCalTotals: number[] = [];
  let mostCals = 0;

  for (let i = 0; i < elves.length; i++) {
    let totalCals = getTotalCals(elves[i]);
    elfCalTotals.push(totalCals);
  }

  elfCalTotals.sort(function (a, b) {
    return b - a;
  });

  mostCals += elfCalTotals[0] + elfCalTotals[1] + elfCalTotals[2];
  console.log(mostCals);
}

main();
