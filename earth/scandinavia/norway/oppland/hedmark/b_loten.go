package hedmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勒滕LotenBarony struct {
	feud.BaseBarony
}

var BLoten勒滕 feud.Barony = &勒滕LotenBarony{}

func init() {
	f := BLoten勒滕.(*勒滕LotenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loten",
		TitleName: "勒滕",
		TitleCode: "b_loten",
	}
}
