package samoyeds

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏哈尼哈SukhanikhaBarony struct {
	feud.BaseBarony
}

var BSukhanikha苏哈尼哈 feud.Barony = &苏哈尼哈SukhanikhaBarony{}

func init() {
    f := BSukhanikha苏哈尼哈.(*苏哈尼哈SukhanikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sukhanikha",
		TitleName: "苏哈尼哈",
		TitleCode: "b_sukhanikha",
	}
}
