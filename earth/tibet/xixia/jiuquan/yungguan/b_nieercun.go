package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 聂耳村NieercunBarony struct {
	feud.BaseBarony
}

var BNieercun聂耳村 feud.Barony = &聂耳村NieercunBarony{}

func init() {
	f := BNieercun聂耳村.(*聂耳村NieercunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nieercun",
		TitleName: "聂耳村",
		TitleCode: "b_nieercun",
	}
}
