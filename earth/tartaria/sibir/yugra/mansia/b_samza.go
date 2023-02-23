package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆扎SamzaBarony struct {
	feud.BaseBarony
}

var BSamza萨姆扎 feud.Barony = &萨姆扎SamzaBarony{}

func init() {
	f := BSamza萨姆扎.(*萨姆扎SamzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samza",
		TitleName: "萨姆扎",
		TitleCode: "b_samza",
	}
}
