package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玛纳斯ManasBarony struct {
	feud.BaseBarony
}

var BManas玛纳斯 feud.Barony = &玛纳斯ManasBarony{}

func init() {
    f := BManas玛纳斯.(*玛纳斯ManasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manas",
		TitleName: "玛纳斯",
		TitleCode: "b_manas",
	}
}
