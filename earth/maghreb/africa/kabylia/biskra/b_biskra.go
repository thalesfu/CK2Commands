package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比斯卡拉BiskraBarony struct {
	feud.BaseBarony
}

var BBiskra比斯卡拉 feud.Barony = &比斯卡拉BiskraBarony{}

func init() {
	f := BBiskra比斯卡拉.(*比斯卡拉BiskraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biskra",
		TitleName: "比斯卡拉",
		TitleCode: "b_biskra",
	}
}
