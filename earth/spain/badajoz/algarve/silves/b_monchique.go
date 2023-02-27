package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙希克MonchiqueBarony struct {
	feud.BaseBarony
}

var BMonchique蒙希克 feud.Barony = &蒙希克MonchiqueBarony{}

func init() {
    f := BMonchique蒙希克.(*蒙希克MonchiqueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monchique",
		TitleName: "蒙希克",
		TitleCode: "b_monchique",
	}
}
