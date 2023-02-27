package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希贝赫HibehBarony struct {
	feud.BaseBarony
}

var BHibeh希贝赫 feud.Barony = &希贝赫HibehBarony{}

func init() {
    f := BHibeh希贝赫.(*希贝赫HibehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hibeh",
		TitleName: "希贝赫",
		TitleCode: "b_hibeh",
	}
}
