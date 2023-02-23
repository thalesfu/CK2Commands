package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 玉普YupuBarony struct {
	feud.BaseBarony
}

var BYupu玉普 feud.Barony = &玉普YupuBarony{}

func init() {
	f := BYupu玉普.(*玉普YupuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yupu",
		TitleName: "玉普",
		TitleCode: "b_yupu",
	}
}
