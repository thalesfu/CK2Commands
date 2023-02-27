package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 湖心坝HuxinbaBarony struct {
	feud.BaseBarony
}

var BHuxinba湖心坝 feud.Barony = &湖心坝HuxinbaBarony{}

func init() {
    f := BHuxinba湖心坝.(*湖心坝HuxinbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huxinba",
		TitleName: "湖心坝",
		TitleCode: "b_huxinba",
	}
}
