package khuttal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙克MunkBarony struct {
	feud.BaseBarony
}

var BMunk蒙克 feud.Barony = &蒙克MunkBarony{}

func init() {
	f := BMunk蒙克.(*蒙克MunkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "munk",
		TitleName: "蒙克",
		TitleCode: "b_munk",
	}
}
