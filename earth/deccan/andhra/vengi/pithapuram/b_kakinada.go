package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡基纳达KakinadaBarony struct {
	feud.BaseBarony
}

var BKakinada卡基纳达 feud.Barony = &卡基纳达KakinadaBarony{}

func init() {
    f := BKakinada卡基纳达.(*卡基纳达KakinadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kakinada",
		TitleName: "卡基纳达",
		TitleCode: "b_kakinada",
	}
}
