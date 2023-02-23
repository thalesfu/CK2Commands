package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅格YagBarony struct {
	feud.BaseBarony
}

var BYag雅格 feud.Barony = &雅格YagBarony{}

func init() {
	f := BYag雅格.(*雅格YagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yag",
		TitleName: "雅格",
		TitleCode: "b_yag",
	}
}
