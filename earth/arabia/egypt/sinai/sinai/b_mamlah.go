package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马穆拉MamlahBarony struct {
	feud.BaseBarony
}

var BMamlah马穆拉 feud.Barony = &马穆拉MamlahBarony{}

func init() {
	f := BMamlah马穆拉.(*马穆拉MamlahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mamlah",
		TitleName: "马穆拉",
		TitleCode: "b_mamlah",
	}
}
