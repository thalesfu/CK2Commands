package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖亚GyaBarony struct {
	feud.BaseBarony
}

var BGya盖亚 feud.Barony = &盖亚GyaBarony{}

func init() {
	f := BGya盖亚.(*盖亚GyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gya",
		TitleName: "盖亚",
		TitleCode: "b_gya",
	}
}
