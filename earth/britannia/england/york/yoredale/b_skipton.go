package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯基普顿SkiptonBarony struct {
	feud.BaseBarony
}

var BSkipton斯基普顿 feud.Barony = &斯基普顿SkiptonBarony{}

func init() {
	f := BSkipton斯基普顿.(*斯基普顿SkiptonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skipton",
		TitleName: "斯基普顿",
		TitleCode: "b_skipton",
	}
}
