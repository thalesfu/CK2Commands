package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 颇罗梨地PhalarddhiBarony struct {
	feud.BaseBarony
}

var BPhalarddhi颇罗梨地 feud.Barony = &颇罗梨地PhalarddhiBarony{}

func init() {
	f := BPhalarddhi颇罗梨地.(*颇罗梨地PhalarddhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "phalarddhi",
		TitleName: "颇罗梨地",
		TitleCode: "b_phalarddhi",
	}
}
