package git

import "os/exec"

func GetUserName() string {
    out, err := exec.Command("git", "config", "user.name").Output()
    if err != nil {
        return ""
    }
    return string(out)
}
