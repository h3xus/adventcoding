const fs = require('fs');

function sumDigits(s) {
    let numz = '';
    let tempArr = [];

    for (let char of s) {
        if (/\d/.test(char)) {
            const digit = parseInt(char);

            if (!isNaN(digit)) {
                tempArr.push(digit);
                numz = tempArr[0] + '' + tempArr[tempArr.length - 1];
            }
        }
    }

    return numz;
}

function main() {
    try {
        const data = fs.readFileSync('../../go/day1/lefile1.txt', 'utf8').split('\n');
        let total = 0;

        for (let line of data) {
            const lineSum = sumDigits(line);

            console.log(`Line: ${line}, Double Number: ${lineSum}`);
            const lineDigit = parseInt(lineSum);

            if (!isNaN(lineDigit)) {
                total += lineDigit;
            }
        }

        console.log(`Total: ${total}`);
    } catch (err) {
        console.error('Error reading file:', err);
    }
}

main();
