package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫瓦尔HvarBarony struct {
	feud.BaseBarony
}

var BHvar赫瓦尔 feud.Barony = &赫瓦尔HvarBarony{}

func init() {
    f := BHvar赫瓦尔.(*赫瓦尔HvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hvar",
		TitleName: "赫瓦尔",
		TitleCode: "b_hvar",
	}
}
