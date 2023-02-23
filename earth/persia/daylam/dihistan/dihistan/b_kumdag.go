package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古姆达格KumdagBarony struct {
	feud.BaseBarony
}

var BKumdag古姆达格 feud.Barony = &古姆达格KumdagBarony{}

func init() {
	f := BKumdag古姆达格.(*古姆达格KumdagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumdag",
		TitleName: "古姆达格",
		TitleCode: "b_kumdag",
	}
}
