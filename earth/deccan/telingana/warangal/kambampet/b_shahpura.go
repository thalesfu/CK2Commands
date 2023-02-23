package kambampet

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙普拉ShahpuraBarony struct {
	feud.BaseBarony
}

var BShahpura沙普拉 feud.Barony = &沙普拉ShahpuraBarony{}

func init() {
	f := BShahpura沙普拉.(*沙普拉ShahpuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahpura",
		TitleName: "沙普拉",
		TitleCode: "b_shahpura",
	}
}
