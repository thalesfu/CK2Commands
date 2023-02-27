package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑格拉SanglaBarony struct {
	feud.BaseBarony
}

var BSangla桑格拉 feud.Barony = &桑格拉SanglaBarony{}

func init() {
    f := BSangla桑格拉.(*桑格拉SanglaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangla",
		TitleName: "桑格拉",
		TitleCode: "b_sangla",
	}
}
