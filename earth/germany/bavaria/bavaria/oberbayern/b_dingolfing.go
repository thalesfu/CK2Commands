package oberbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 丁戈尔芬DingolfingBarony struct {
	feud.BaseBarony
}

var BDingolfing丁戈尔芬 feud.Barony = &丁戈尔芬DingolfingBarony{}

func init() {
    f := BDingolfing丁戈尔芬.(*丁戈尔芬DingolfingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dingolfing",
		TitleName: "丁戈尔芬",
		TitleCode: "b_dingolfing",
	}
}
