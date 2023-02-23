package ikonion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉兰达LarandaBarony struct {
	feud.BaseBarony
}

var BLaranda拉兰达 feud.Barony = &拉兰达LarandaBarony{}

func init() {
	f := BLaranda拉兰达.(*拉兰达LarandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laranda",
		TitleName: "拉兰达",
		TitleCode: "b_laranda",
	}
}
