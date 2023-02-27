package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉加德RaigadBarony struct {
	feud.BaseBarony
}

var BRaigad拉加德 feud.Barony = &拉加德RaigadBarony{}

func init() {
    f := BRaigad拉加德.(*拉加德RaigadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raigad",
		TitleName: "拉加德",
		TitleCode: "b_raigad",
	}
}
