package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷维利奥TreviglioBarony struct {
	feud.BaseBarony
}

var BTreviglio特雷维利奥 feud.Barony = &特雷维利奥TreviglioBarony{}

func init() {
    f := BTreviglio特雷维利奥.(*特雷维利奥TreviglioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treviglio",
		TitleName: "特雷维利奥",
		TitleCode: "b_treviglio",
	}
}
