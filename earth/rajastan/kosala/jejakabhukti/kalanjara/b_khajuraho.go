package kalanjara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 渴树罗城KhajurahoBarony struct {
	feud.BaseBarony
}

var BKhajuraho渴树罗城 feud.Barony = &渴树罗城KhajurahoBarony{}

func init() {
    f := BKhajuraho渴树罗城.(*渴树罗城KhajurahoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khajuraho",
		TitleName: "渴树罗城",
		TitleCode: "b_khajuraho",
	}
}
