package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 石屯ShitunBarony struct {
	feud.BaseBarony
}

var BShitun石屯 feud.Barony = &石屯ShitunBarony{}

func init() {
	f := BShitun石屯.(*石屯ShitunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shitun",
		TitleName: "石屯",
		TitleCode: "b_shitun",
	}
}
