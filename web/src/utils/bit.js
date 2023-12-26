
export const toBitInt = (arr) => {
    const max = Math.max(...arr);
    const newArr = new Array(max+1).fill(0);
    for (let i = 0; i < arr.length; i++) {
        newArr[arr[i]] = 1;
    }
    let num = 0b0;
    for (let i = 0; i < newArr.length; i++) {
        num = num << 1;
        num = num + newArr[i];
    }
    return num;
}

export const toArrFromBitInt = (num) => {
    const arr = []; 
    while (num > 0) { 
        arr.unshift(num % 2); 
        num = Math.floor(num / 2); 
    } 
    const newArr = []; 
    for (let i = 0; i < arr.length; i++) {
         if (arr[i] !== 0) { 
            newArr.push(i+1); 
        } 
    } 
    return newArr.length === 0 ? [0] : newArr; 
}