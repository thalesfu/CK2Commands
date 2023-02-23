package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀拉特斯吉格特QalatsjergatBarony struct {
	feud.BaseBarony
}

var BQalatsjergat喀拉特斯吉格特 feud.Barony = &喀拉特斯吉格特QalatsjergatBarony{}

func init() {
	f := BQalatsjergat喀拉特斯吉格特.(*喀拉特斯吉格特QalatsjergatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qalatsjergat",
		TitleName: "喀拉特斯吉格特",
		TitleCode: "b_qalatsjergat",
	}
}
