package gelre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈梅亨NijmegenBarony struct {
	feud.BaseBarony
}

var BNijmegen奈梅亨 feud.Barony = &奈梅亨NijmegenBarony{}

func init() {
    f := BNijmegen奈梅亨.(*奈梅亨NijmegenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nijmegen",
		TitleName: "奈梅亨",
		TitleCode: "b_nijmegen",
	}
}
