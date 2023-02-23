package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙哥马利MontgomeryBarony struct {
	feud.BaseBarony
}

var BMontgomery蒙哥马利 feud.Barony = &蒙哥马利MontgomeryBarony{}

func init() {
	f := BMontgomery蒙哥马利.(*蒙哥马利MontgomeryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montgomery",
		TitleName: "蒙哥马利",
		TitleCode: "b_montgomery",
	}
}
