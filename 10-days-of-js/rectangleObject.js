function Rectangle(a, b) {
    const perimeter = 2 * (a + b)
    const area = a * b
    return {length: a, width: b, perimeter, area}
}

const rec = new Rectangle(10, 9)

console.log(rec.length)