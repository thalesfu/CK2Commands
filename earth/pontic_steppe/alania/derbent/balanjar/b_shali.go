package balanjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙利ShaliBarony struct {
	feud.BaseBarony
}

var BShali沙利 feud.Barony = &沙利ShaliBarony{}

func init() {
    f := BShali沙利.(*沙利ShaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shali",
		TitleName: "沙利",
		TitleCode: "b_shali",
	}
}
