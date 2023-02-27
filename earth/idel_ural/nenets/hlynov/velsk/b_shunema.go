package velsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒涅马ShunemaBarony struct {
	feud.BaseBarony
}

var BShunema舒涅马 feud.Barony = &舒涅马ShunemaBarony{}

func init() {
    f := BShunema舒涅马.(*舒涅马ShunemaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shunema",
		TitleName: "舒涅马",
		TitleCode: "b_shunema",
	}
}
