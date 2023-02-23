package lenghu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 育才YucaiBarony struct {
	feud.BaseBarony
}

var BYucai育才 feud.Barony = &育才YucaiBarony{}

func init() {
	f := BYucai育才.(*育才YucaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yucai",
		TitleName: "育才",
		TitleCode: "b_yucai",
	}
}
