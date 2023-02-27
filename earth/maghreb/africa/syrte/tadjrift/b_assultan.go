package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏丹AssultanBarony struct {
	feud.BaseBarony
}

var BAssultan苏丹 feud.Barony = &苏丹AssultanBarony{}

func init() {
    f := BAssultan苏丹.(*苏丹AssultanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assultan",
		TitleName: "苏丹",
		TitleCode: "b_assultan",
	}
}
