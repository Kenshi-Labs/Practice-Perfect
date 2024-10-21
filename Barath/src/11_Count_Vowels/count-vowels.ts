export function countVowels(str: string): number {
    const vowels = ['a', 'e', 'i', 'o', 'u'];
    const vowelCount = str
        .toLowerCase()
        .split('')
        .filter(char => vowels.includes(char)).length;

    console.log(`Number of vowels in "${str}":`, vowelCount);  

    return vowelCount;
}
