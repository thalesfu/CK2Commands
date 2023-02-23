package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 绍阿ShewaBarony struct {
	feud.BaseBarony
}

var BShewa绍阿 feud.Barony = &绍阿ShewaBarony{}

func init() {
	f := BShewa绍阿.(*绍阿ShewaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shewa",
		TitleName: "绍阿",
		TitleCode: "b_shewa",
	}
}
