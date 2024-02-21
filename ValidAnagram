func isAnagram(s string, t string) bool {
    Amap := make(map[rune]int)

    for _, v := range s {
        Amap[v]++
    }

    for _, v := range t{
        Amap[v]--
    }

    for _, value := range Amap {
        if value != 0 {
            return false
        }
    }

    return true
}
