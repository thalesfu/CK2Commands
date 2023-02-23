package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 竹根滩ZhugentanBarony struct {
	feud.BaseBarony
}

var BZhugentan竹根滩 feud.Barony = &竹根滩ZhugentanBarony{}

func init() {
	f := BZhugentan竹根滩.(*竹根滩ZhugentanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhugentan",
		TitleName: "竹根滩",
		TitleCode: "b_zhugentan",
	}
}
