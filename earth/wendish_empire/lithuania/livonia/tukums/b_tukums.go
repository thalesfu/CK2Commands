package tukums

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图库姆斯TukumsBarony struct {
	feud.BaseBarony
}

var BTukums图库姆斯 feud.Barony = &图库姆斯TukumsBarony{}

func init() {
    f := BTukums图库姆斯.(*图库姆斯TukumsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tukums",
		TitleName: "图库姆斯",
		TitleCode: "b_tukums",
	}
}
