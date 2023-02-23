package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛永SeiyunBarony struct {
	feud.BaseBarony
}

var BSeiyun赛永 feud.Barony = &赛永SeiyunBarony{}

func init() {
	f := BSeiyun赛永.(*赛永SeiyunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seiyun",
		TitleName: "赛永",
		TitleCode: "b_seiyun",
	}
}
