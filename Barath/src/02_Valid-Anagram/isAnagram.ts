function isAnagram(s: string, t: string): boolean {
    if (s.length !== t.length) {
        return false;
    }

    const sortString = (str: string) => str.split('').sort().join('');
    
    return sortString(s) === sortString(t);
}


console.log(isAnagram('listen', 'silent')); 

export default isAnagram;