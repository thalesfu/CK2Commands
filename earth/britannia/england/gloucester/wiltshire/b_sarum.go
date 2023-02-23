package wiltshire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞勒姆SarumBarony struct {
	feud.BaseBarony
}

var BSarum塞勒姆 feud.Barony = &塞勒姆SarumBarony{}

func init() {
	f := BSarum塞勒姆.(*塞勒姆SarumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarum",
		TitleName: "塞勒姆",
		TitleCode: "b_sarum",
	}
}
