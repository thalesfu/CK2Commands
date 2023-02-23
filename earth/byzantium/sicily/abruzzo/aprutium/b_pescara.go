package aprutium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩斯卡拉PescaraBarony struct {
	feud.BaseBarony
}

var BPescara佩斯卡拉 feud.Barony = &佩斯卡拉PescaraBarony{}

func init() {
	f := BPescara佩斯卡拉.(*佩斯卡拉PescaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pescara",
		TitleName: "佩斯卡拉",
		TitleCode: "b_pescara",
	}
}
