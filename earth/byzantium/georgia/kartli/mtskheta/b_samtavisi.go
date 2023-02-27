package mtskheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆塔维西SamtavisiBarony struct {
	feud.BaseBarony
}

var BSamtavisi萨姆塔维西 feud.Barony = &萨姆塔维西SamtavisiBarony{}

func init() {
    f := BSamtavisi萨姆塔维西.(*萨姆塔维西SamtavisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samtavisi",
		TitleName: "萨姆塔维西",
		TitleCode: "b_samtavisi",
	}
}
