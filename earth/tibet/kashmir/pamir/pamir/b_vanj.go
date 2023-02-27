package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 万吉VanjBarony struct {
	feud.BaseBarony
}

var BVanj万吉 feud.Barony = &万吉VanjBarony{}

func init() {
    f := BVanj万吉.(*万吉VanjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vanj",
		TitleName: "万吉",
		TitleCode: "b_vanj",
	}
}
