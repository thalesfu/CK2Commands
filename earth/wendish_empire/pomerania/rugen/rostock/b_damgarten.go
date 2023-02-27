package rostock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达姆加滕DamgartenBarony struct {
	feud.BaseBarony
}

var BDamgarten达姆加滕 feud.Barony = &达姆加滕DamgartenBarony{}

func init() {
    f := BDamgarten达姆加滕.(*达姆加滕DamgartenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "damgarten",
		TitleName: "达姆加滕",
		TitleCode: "b_damgarten",
	}
}
