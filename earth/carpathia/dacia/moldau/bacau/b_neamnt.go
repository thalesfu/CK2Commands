package bacau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼亚姆茨NeamntBarony struct {
	feud.BaseBarony
}

var BNeamnt尼亚姆茨 feud.Barony = &尼亚姆茨NeamntBarony{}

func init() {
	f := BNeamnt尼亚姆茨.(*尼亚姆茨NeamntBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neamnt",
		TitleName: "尼亚姆茨",
		TitleCode: "b_neamnt",
	}
}
