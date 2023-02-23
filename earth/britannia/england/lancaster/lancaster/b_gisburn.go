package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉斯本GisburnBarony struct {
	feud.BaseBarony
}

var BGisburn吉斯本 feud.Barony = &吉斯本GisburnBarony{}

func init() {
	f := BGisburn吉斯本.(*吉斯本GisburnBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gisburn",
		TitleName: "吉斯本",
		TitleCode: "b_gisburn",
	}
}
