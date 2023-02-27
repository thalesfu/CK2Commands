package kebbi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯比KebbiBarony struct {
	feud.BaseBarony
}

var BKebbi凯比 feud.Barony = &凯比KebbiBarony{}

func init() {
    f := BKebbi凯比.(*凯比KebbiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kebbi",
		TitleName: "凯比",
		TitleCode: "b_kebbi",
	}
}
