package tobruk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜尔迪耶BardiaBarony struct {
	feud.BaseBarony
}

var BBardia拜尔迪耶 feud.Barony = &拜尔迪耶BardiaBarony{}

func init() {
	f := BBardia拜尔迪耶.(*拜尔迪耶BardiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bardia",
		TitleName: "拜尔迪耶",
		TitleCode: "b_bardia",
	}
}
