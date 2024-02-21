func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string)
    for _, str := range strs {
        sortedStr := sortString(str)
        anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
    }

    result := make([][]string, 0)

    // Проходим по значениям карты и добавляем их в результат
    for _, v := range anagramMap {
        result = append(result, v)
    }

    return result
}

func sortString(s string) string {
    chars := []byte(s)
    sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
    return string(chars)
}
