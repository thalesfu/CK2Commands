package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博讷BeauneBarony struct {
	feud.BaseBarony
}

var BBeaune博讷 feud.Barony = &博讷BeauneBarony{}

func init() {
	f := BBeaune博讷.(*博讷BeauneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaune",
		TitleName: "博讷",
		TitleCode: "b_beaune",
	}
}
