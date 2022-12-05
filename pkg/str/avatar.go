package str

import "fmt"

func GenerateRandomAvatar(seed string) string {
	return fmt.Sprintf("https://avatars.dicebear.com/api/identicon/%s.svg?size=256", seed)
}
