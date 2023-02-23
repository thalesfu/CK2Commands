package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朗盖LangayBarony struct {
	feud.BaseBarony
}

var BLangay朗盖 feud.Barony = &朗盖LangayBarony{}

func init() {
	f := BLangay朗盖.(*朗盖LangayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "langay",
		TitleName: "朗盖",
		TitleCode: "b_langay",
	}
}
