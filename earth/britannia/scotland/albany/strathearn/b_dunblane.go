package strathearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓布兰DunblaneBarony struct {
	feud.BaseBarony
}

var BDunblane邓布兰 feud.Barony = &邓布兰DunblaneBarony{}

func init() {
	f := BDunblane邓布兰.(*邓布兰DunblaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunblane",
		TitleName: "邓布兰",
		TitleCode: "b_dunblane",
	}
}
