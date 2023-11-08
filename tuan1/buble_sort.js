const bubbleSort = (array) => {
    let isOrdered;
    for (let i = 0; i < array.length; i++) {
        isOrdered = true; // biến để kiếm tra vòng lặp và break
        for (let x = 0; x < array.length - 1 - i; x++) {
            if (array[x] > array[x + 1]) {
                [array[x], array[x + 1]] = [array[x + 1], array[x]];
                isOrdered = false;
            }
        }
        if (isOrdered) break;
    }
    return array;
}

let a = [3, 1, 2]
console.log(bubbleSort(a))
