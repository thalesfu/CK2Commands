package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗滕堡Rotenburg_hesseBarony struct {
	feud.BaseBarony
}

var BRotenburg_hesse罗滕堡 feud.Barony = &罗滕堡Rotenburg_hesseBarony{}

func init() {
    f := BRotenburg_hesse罗滕堡.(*罗滕堡Rotenburg_hesseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rotenburg_hesse",
		TitleName: "罗滕堡",
		TitleCode: "b_rotenburg_hesse",
	}
}
