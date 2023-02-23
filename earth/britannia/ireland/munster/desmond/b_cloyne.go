package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克洛因CloyneBarony struct {
	feud.BaseBarony
}

var BCloyne克洛因 feud.Barony = &克洛因CloyneBarony{}

func init() {
	f := BCloyne克洛因.(*克洛因CloyneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cloyne",
		TitleName: "克洛因",
		TitleCode: "b_cloyne",
	}
}
