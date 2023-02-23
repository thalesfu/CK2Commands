package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉克KerakBarony struct {
	feud.BaseBarony
}

var BKerak克拉克 feud.Barony = &克拉克KerakBarony{}

func init() {
	f := BKerak克拉克.(*克拉克KerakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerak",
		TitleName: "克拉克",
		TitleCode: "b_kerak",
	}
}
