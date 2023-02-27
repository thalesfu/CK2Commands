package lattalura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗多楼罗LattaluraBarony struct {
	feud.BaseBarony
}

var BLattalura罗多楼罗 feud.Barony = &罗多楼罗LattaluraBarony{}

func init() {
    f := BLattalura罗多楼罗.(*罗多楼罗LattaluraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lattalura",
		TitleName: "罗多楼罗",
		TitleCode: "b_lattalura",
	}
}
