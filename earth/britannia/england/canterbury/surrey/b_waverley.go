package surrey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦弗利WaverleyBarony struct {
	feud.BaseBarony
}

var BWaverley韦弗利 feud.Barony = &韦弗利WaverleyBarony{}

func init() {
	f := BWaverley韦弗利.(*韦弗利WaverleyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waverley",
		TitleName: "韦弗利",
		TitleCode: "b_waverley",
	}
}
