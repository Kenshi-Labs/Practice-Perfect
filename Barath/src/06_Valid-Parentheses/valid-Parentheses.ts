function isValid(s: string): boolean {
    const stack: string[] = [];

    for (let char of s) {
        if (char === '(' || char === '{' || char === '[') {
            stack.push(char);
        } else {
            const top = stack.pop();
            if (
                (char === ')' && top !== '(') ||
                (char === '}' && top !== '{') ||
                (char === ']' && top !== '[')
            ) {
                return false;
            }
        }
    }

    return stack.length === 0;
}

console.log(isValid("()[]{}")); 

export default isValid;
