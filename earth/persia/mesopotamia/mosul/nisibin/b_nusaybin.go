package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努赛宾NusaybinBarony struct {
	feud.BaseBarony
}

var BNusaybin努赛宾 feud.Barony = &努赛宾NusaybinBarony{}

func init() {
    f := BNusaybin努赛宾.(*努赛宾NusaybinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nusaybin",
		TitleName: "努赛宾",
		TitleCode: "b_nusaybin",
	}
}
