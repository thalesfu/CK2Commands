package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔斯克斯肯TaskeskenBarony struct {
	feud.BaseBarony
}

var BTaskesken塔斯克斯肯 feud.Barony = &塔斯克斯肯TaskeskenBarony{}

func init() {
	f := BTaskesken塔斯克斯肯.(*塔斯克斯肯TaskeskenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taskesken",
		TitleName: "塔斯克斯肯",
		TitleCode: "b_taskesken",
	}
}
