package galloway

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓拉吉特DunragitBarony struct {
	feud.BaseBarony
}

var BDunragit邓拉吉特 feud.Barony = &邓拉吉特DunragitBarony{}

func init() {
    f := BDunragit邓拉吉特.(*邓拉吉特DunragitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunragit",
		TitleName: "邓拉吉特",
		TitleCode: "b_dunragit",
	}
}
