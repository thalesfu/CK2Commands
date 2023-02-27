package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉克塔RakteBarony struct {
	feud.BaseBarony
}

var BRakte拉克塔 feud.Barony = &拉克塔RakteBarony{}

func init() {
    f := BRakte拉克塔.(*拉克塔RakteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rakte",
		TitleName: "拉克塔",
		TitleCode: "b_rakte",
	}
}
