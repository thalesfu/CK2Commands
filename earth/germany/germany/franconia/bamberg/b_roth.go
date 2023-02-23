package bamberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗特RothBarony struct {
	feud.BaseBarony
}

var BRoth罗特 feud.Barony = &罗特RothBarony{}

func init() {
	f := BRoth罗特.(*罗特RothBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roth",
		TitleName: "罗特",
		TitleCode: "b_roth",
	}
}
