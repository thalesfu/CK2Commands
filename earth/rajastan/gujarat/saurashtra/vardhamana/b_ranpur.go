package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰补罗RanpurBarony struct {
	feud.BaseBarony
}

var BRanpur兰补罗 feud.Barony = &兰补罗RanpurBarony{}

func init() {
    f := BRanpur兰补罗.(*兰补罗RanpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ranpur",
		TitleName: "兰补罗",
		TitleCode: "b_ranpur",
	}
}
