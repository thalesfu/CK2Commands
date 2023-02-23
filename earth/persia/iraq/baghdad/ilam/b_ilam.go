package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊拉姆IlamBarony struct {
	feud.BaseBarony
}

var BIlam伊拉姆 feud.Barony = &伊拉姆IlamBarony{}

func init() {
	f := BIlam伊拉姆.(*伊拉姆IlamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ilam",
		TitleName: "伊拉姆",
		TitleCode: "b_ilam",
	}
}
