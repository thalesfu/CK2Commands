package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赤麦干ShymkentBarony struct {
	feud.BaseBarony
}

var BShymkent赤麦干 feud.Barony = &赤麦干ShymkentBarony{}

func init() {
    f := BShymkent赤麦干.(*赤麦干ShymkentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shymkent",
		TitleName: "赤麦干",
		TitleCode: "b_shymkent",
	}
}
