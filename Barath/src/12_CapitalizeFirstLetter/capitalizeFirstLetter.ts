// capitalizeFirstLetter.ts

export function capitalizeFirstLetter(str: string): string {
    const result = str
        .split(' ')  
        .map(word => word.charAt(0).toUpperCase() + word.slice(1))  
        .join(' ');  
    
    console.log(`Capitalized string: "${result}"`);

    return result;
}

const input = "the quick brown fox";
capitalizeFirstLetter(input);  
