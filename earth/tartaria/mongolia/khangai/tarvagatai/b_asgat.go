package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯嘎特AsgatBarony struct {
	feud.BaseBarony
}

var BAsgat阿斯嘎特 feud.Barony = &阿斯嘎特AsgatBarony{}

func init() {
    f := BAsgat阿斯嘎特.(*阿斯嘎特AsgatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asgat",
		TitleName: "阿斯嘎特",
		TitleCode: "b_asgat",
	}
}
