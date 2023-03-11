package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克鲁佩茨KrupetsBarony struct {
	feud.BaseBarony
}

var BKrupets克鲁佩茨 feud.Barony = &克鲁佩茨KrupetsBarony{}

func init() {
    f := BKrupets克鲁佩茨.(*克鲁佩茨KrupetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krupets",
		TitleName: "克鲁佩茨",
		TitleCode: "b_krupets",
	}
}
