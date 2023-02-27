package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久舍KyusheBarony struct {
	feud.BaseBarony
}

var BKyushe久舍 feud.Barony = &久舍KyusheBarony{}

func init() {
    f := BKyushe久舍.(*久舍KyusheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyushe",
		TitleName: "久舍",
		TitleCode: "b_kyushe",
	}
}
