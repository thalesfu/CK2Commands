package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马希纳MashinaBarony struct {
	feud.BaseBarony
}

var BMashina马希纳 feud.Barony = &马希纳MashinaBarony{}

func init() {
	f := BMashina马希纳.(*马希纳MashinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mashina",
		TitleName: "马希纳",
		TitleCode: "b_mashina",
	}
}
