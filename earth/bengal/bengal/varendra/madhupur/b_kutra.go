package madhupur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠多罗KutraBarony struct {
	feud.BaseBarony
}

var BKutra鸠多罗 feud.Barony = &鸠多罗KutraBarony{}

func init() {
	f := BKutra鸠多罗.(*鸠多罗KutraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kutra",
		TitleName: "鸠多罗",
		TitleCode: "b_kutra",
	}
}
