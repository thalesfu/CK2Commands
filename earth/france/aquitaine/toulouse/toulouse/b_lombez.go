package toulouse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆贝LombezBarony struct {
	feud.BaseBarony
}

var BLombez隆贝 feud.Barony = &隆贝LombezBarony{}

func init() {
	f := BLombez隆贝.(*隆贝LombezBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lombez",
		TitleName: "隆贝",
		TitleCode: "b_lombez",
	}
}
