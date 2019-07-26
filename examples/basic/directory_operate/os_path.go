package directory_operate

import (
    "os"
    "os/user"
)

func DirOperations() (string, error) {
    var dir string

    u, err := user.Current()
    if err == nil {
        dir = u.HomeDir
    } else if home := os.Getenv("HOME"); home != "" {
        dir = home
    } else {
        wd, err := os.Getwd()
        if err != nil {
            return "", err
        }
        dir = wd
    }

    return dir, nil
}
