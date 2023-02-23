package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾雷拉JararaBarony struct {
	feud.BaseBarony
}

var BJarara贾雷拉 feud.Barony = &贾雷拉JararaBarony{}

func init() {
	f := BJarara贾雷拉.(*贾雷拉JararaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarara",
		TitleName: "贾雷拉",
		TitleCode: "b_jarara",
	}
}
