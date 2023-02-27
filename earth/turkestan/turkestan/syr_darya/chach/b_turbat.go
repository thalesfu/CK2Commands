package chach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 土尔巴特TurbatBarony struct {
	feud.BaseBarony
}

var BTurbat土尔巴特 feud.Barony = &土尔巴特TurbatBarony{}

func init() {
    f := BTurbat土尔巴特.(*土尔巴特TurbatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turbat",
		TitleName: "土尔巴特",
		TitleCode: "b_turbat",
	}
}
