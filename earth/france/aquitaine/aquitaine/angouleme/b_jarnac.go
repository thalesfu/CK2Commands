package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅尔纳克JarnacBarony struct {
	feud.BaseBarony
}

var BJarnac雅尔纳克 feud.Barony = &雅尔纳克JarnacBarony{}

func init() {
	f := BJarnac雅尔纳克.(*雅尔纳克JarnacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarnac",
		TitleName: "雅尔纳克",
		TitleCode: "b_jarnac",
	}
}
