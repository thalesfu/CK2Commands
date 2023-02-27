package verden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗滕堡Rotenburg_verdenBarony struct {
	feud.BaseBarony
}

var BRotenburg_verden罗滕堡 feud.Barony = &罗滕堡Rotenburg_verdenBarony{}

func init() {
    f := BRotenburg_verden罗滕堡.(*罗滕堡Rotenburg_verdenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rotenburg_verden",
		TitleName: "罗滕堡",
		TitleCode: "b_rotenburg_verden",
	}
}
