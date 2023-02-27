package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达拉喀TeglakarBarony struct {
	feud.BaseBarony
}

var BTeglakar达拉喀 feud.Barony = &达拉喀TeglakarBarony{}

func init() {
    f := BTeglakar达拉喀.(*达拉喀TeglakarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teglakar",
		TitleName: "达拉喀",
		TitleCode: "b_teglakar",
	}
}
