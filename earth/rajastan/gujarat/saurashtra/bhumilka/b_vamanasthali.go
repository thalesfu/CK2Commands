package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 筏摩那悉他利VamanasthaliBarony struct {
	feud.BaseBarony
}

var BVamanasthali筏摩那悉他利 feud.Barony = &筏摩那悉他利VamanasthaliBarony{}

func init() {
    f := BVamanasthali筏摩那悉他利.(*筏摩那悉他利VamanasthaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vamanasthali",
		TitleName: "筏摩那悉他利",
		TitleCode: "b_vamanasthali",
	}
}
