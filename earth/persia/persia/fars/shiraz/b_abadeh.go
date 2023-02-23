package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴代AbadehBarony struct {
	feud.BaseBarony
}

var BAbadeh阿巴代 feud.Barony = &阿巴代AbadehBarony{}

func init() {
	f := BAbadeh阿巴代.(*阿巴代AbadehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abadeh",
		TitleName: "阿巴代",
		TitleCode: "b_abadeh",
	}
}
