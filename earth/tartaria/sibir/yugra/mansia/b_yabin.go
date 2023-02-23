package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚宾YabinBarony struct {
	feud.BaseBarony
}

var BYabin亚宾 feud.Barony = &亚宾YabinBarony{}

func init() {
	f := BYabin亚宾.(*亚宾YabinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yabin",
		TitleName: "亚宾",
		TitleCode: "b_yabin",
	}
}
