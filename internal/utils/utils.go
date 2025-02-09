package utils

import "strconv"

func IsValidYear(year string) bool {
    _, err := strconv.Atoi(year)
    return err == nil
}