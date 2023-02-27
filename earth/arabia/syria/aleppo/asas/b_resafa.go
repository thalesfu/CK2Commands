package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁萨费ResafaBarony struct {
	feud.BaseBarony
}

var BResafa鲁萨费 feud.Barony = &鲁萨费ResafaBarony{}

func init() {
    f := BResafa鲁萨费.(*鲁萨费ResafaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "resafa",
		TitleName: "鲁萨费",
		TitleCode: "b_resafa",
	}
}
