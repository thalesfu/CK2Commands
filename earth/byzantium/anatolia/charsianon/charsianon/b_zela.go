package charsianon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽拉ZelaBarony struct {
	feud.BaseBarony
}

var BZela泽拉 feud.Barony = &泽拉ZelaBarony{}

func init() {
	f := BZela泽拉.(*泽拉ZelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zela",
		TitleName: "泽拉",
		TitleCode: "b_zela",
	}
}
