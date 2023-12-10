import { readFile } from 'node:fs';

function main() {
    readFile('input', (err, data) => {
        console.log(data.toString())
    })
}

main()
