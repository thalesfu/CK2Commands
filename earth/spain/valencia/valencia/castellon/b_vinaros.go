package castellon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比纳罗斯VinarosBarony struct {
	feud.BaseBarony
}

var BVinaros比纳罗斯 feud.Barony = &比纳罗斯VinarosBarony{}

func init() {
	f := BVinaros比纳罗斯.(*比纳罗斯VinarosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vinaros",
		TitleName: "比纳罗斯",
		TitleCode: "b_vinaros",
	}
}
