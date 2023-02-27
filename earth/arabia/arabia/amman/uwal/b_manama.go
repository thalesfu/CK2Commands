package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦纳麦ManamaBarony struct {
	feud.BaseBarony
}

var BManama麦纳麦 feud.Barony = &麦纳麦ManamaBarony{}

func init() {
    f := BManama麦纳麦.(*麦纳麦ManamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manama",
		TitleName: "麦纳麦",
		TitleCode: "b_manama",
	}
}
