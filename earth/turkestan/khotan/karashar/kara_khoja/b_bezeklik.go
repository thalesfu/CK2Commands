package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柏孜克里克BezeklikBarony struct {
	feud.BaseBarony
}

var BBezeklik柏孜克里克 feud.Barony = &柏孜克里克BezeklikBarony{}

func init() {
    f := BBezeklik柏孜克里克.(*柏孜克里克BezeklikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bezeklik",
		TitleName: "柏孜克里克",
		TitleCode: "b_bezeklik",
	}
}
