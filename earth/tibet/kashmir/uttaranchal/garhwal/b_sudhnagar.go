package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏提那迦利SudhnagarBarony struct {
	feud.BaseBarony
}

var BSudhnagar苏提那迦利 feud.Barony = &苏提那迦利SudhnagarBarony{}

func init() {
    f := BSudhnagar苏提那迦利.(*苏提那迦利SudhnagarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sudhnagar",
		TitleName: "苏提那迦利",
		TitleCode: "b_sudhnagar",
	}
}
