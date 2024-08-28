package fs

import (
	"fmt"
	"os"
	"os/exec"
)

func Reader(banner string) string {
	// if reading the bannner fails or the file got altered in a way that its gonna ruin the ascii-art it will be  fetched and  that will secure normal running of the program
	content, err := os.ReadFile(banner + ".txt")
	if err != nil || len(content) != 6623 {
		url := "https://learn.zone01oujda.ma/git/root/public/raw/branch/master/subjects/ascii-art/" + banner + ".txt"
		cmd := exec.Command("curl", "-o", "ascii-art/banners/"+banner+".txt", url)
		_, er := cmd.Output()
		if er != nil {
			fmt.Println("Error fetching url:", er)
			os.Exit(1)
		}
		content, _ = os.ReadFile("ascii-art/banners/"+banner + ".txt")
	}
	return string(content)
}
