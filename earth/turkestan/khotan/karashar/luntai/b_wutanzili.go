package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌贪訾离WutanziliBarony struct {
	feud.BaseBarony
}

var BWutanzili乌贪訾离 feud.Barony = &乌贪訾离WutanziliBarony{}

func init() {
	f := BWutanzili乌贪訾离.(*乌贪訾离WutanziliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wutanzili",
		TitleName: "乌贪訾离",
		TitleCode: "b_wutanzili",
	}
}
