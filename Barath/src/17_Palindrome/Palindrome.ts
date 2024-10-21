export function isPalindrome(input: string): boolean {
    const sanitizedInput = input.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();

    return sanitizedInput === sanitizedInput.split('').reverse().join('');
}

console.log(isPalindrome('A man, a plan, a canal: Panama')); 