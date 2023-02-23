package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西格纳吉SighnaghiBarony struct {
	feud.BaseBarony
}

var BSighnaghi西格纳吉 feud.Barony = &西格纳吉SighnaghiBarony{}

func init() {
	f := BSighnaghi西格纳吉.(*西格纳吉SighnaghiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sighnaghi",
		TitleName: "西格纳吉",
		TitleCode: "b_sighnaghi",
	}
}
