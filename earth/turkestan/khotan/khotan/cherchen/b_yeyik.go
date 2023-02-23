package cherchen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶亦克YeyikBarony struct {
	feud.BaseBarony
}

var BYeyik叶亦克 feud.Barony = &叶亦克YeyikBarony{}

func init() {
	f := BYeyik叶亦克.(*叶亦克YeyikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yeyik",
		TitleName: "叶亦克",
		TitleCode: "b_yeyik",
	}
}
