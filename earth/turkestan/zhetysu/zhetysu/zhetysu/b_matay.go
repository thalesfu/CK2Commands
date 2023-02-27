package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马泰MatayBarony struct {
	feud.BaseBarony
}

var BMatay马泰 feud.Barony = &马泰MatayBarony{}

func init() {
    f := BMatay马泰.(*马泰MatayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "matay",
		TitleName: "马泰",
		TitleCode: "b_matay",
	}
}
