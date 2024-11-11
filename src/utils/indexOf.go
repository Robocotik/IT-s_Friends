package utils

func IndexOf(arr []string, need string) int {
    for ind, val := range arr {
        if val == need {
            return ind
        }
    }
    return -1
}