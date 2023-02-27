package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 慕斯MushBarony struct {
	feud.BaseBarony
}

var BMush慕斯 feud.Barony = &慕斯MushBarony{}

func init() {
    f := BMush慕斯.(*慕斯MushBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mush",
		TitleName: "慕斯",
		TitleCode: "b_mush",
	}
}
