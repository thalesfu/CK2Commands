package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏吉SoughiBarony struct {
	feud.BaseBarony
}

var BSoughi苏吉 feud.Barony = &苏吉SoughiBarony{}

func init() {
	f := BSoughi苏吉.(*苏吉SoughiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soughi",
		TitleName: "苏吉",
		TitleCode: "b_soughi",
	}
}
