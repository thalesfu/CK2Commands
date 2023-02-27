package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加拉CaraoasisBarony struct {
	feud.BaseBarony
}

var BCaraoasis加拉 feud.Barony = &加拉CaraoasisBarony{}

func init() {
    f := BCaraoasis加拉.(*加拉CaraoasisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caraoasis",
		TitleName: "加拉",
		TitleCode: "b_caraoasis",
	}
}
