package tirgoviste

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔戈维什泰TargovisteBarony struct {
	feud.BaseBarony
}

var BTargoviste特尔戈维什泰 feud.Barony = &特尔戈维什泰TargovisteBarony{}

func init() {
    f := BTargoviste特尔戈维什泰.(*特尔戈维什泰TargovisteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "targoviste",
		TitleName: "特尔戈维什泰",
		TitleCode: "b_targoviste",
	}
}
