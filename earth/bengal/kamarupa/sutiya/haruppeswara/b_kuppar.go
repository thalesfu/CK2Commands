package haruppeswara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠波KupparBarony struct {
	feud.BaseBarony
}

var BKuppar鸠波 feud.Barony = &鸠波KupparBarony{}

func init() {
	f := BKuppar鸠波.(*鸠波KupparBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuppar",
		TitleName: "鸠波",
		TitleCode: "b_kuppar",
	}
}
