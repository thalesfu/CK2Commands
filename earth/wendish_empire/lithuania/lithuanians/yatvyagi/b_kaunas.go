package yatvyagi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考纳斯KaunasBarony struct {
	feud.BaseBarony
}

var BKaunas考纳斯 feud.Barony = &考纳斯KaunasBarony{}

func init() {
    f := BKaunas考纳斯.(*考纳斯KaunasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaunas",
		TitleName: "考纳斯",
		TitleCode: "b_kaunas",
	}
}
