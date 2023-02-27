package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁姆洛夫KrumlovBarony struct {
	feud.BaseBarony
}

var BKrumlov克鲁姆洛夫 feud.Barony = &克鲁姆洛夫KrumlovBarony{}

func init() {
    f := BKrumlov克鲁姆洛夫.(*克鲁姆洛夫KrumlovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krumlov",
		TitleName: "克鲁姆洛夫",
		TitleCode: "b_krumlov",
	}
}
