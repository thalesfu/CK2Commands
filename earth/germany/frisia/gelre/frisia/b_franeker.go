package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉讷克FranekerBarony struct {
	feud.BaseBarony
}

var BFraneker弗拉讷克 feud.Barony = &弗拉讷克FranekerBarony{}

func init() {
    f := BFraneker弗拉讷克.(*弗拉讷克FranekerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "franeker",
		TitleName: "弗拉讷克",
		TitleCode: "b_franeker",
	}
}
