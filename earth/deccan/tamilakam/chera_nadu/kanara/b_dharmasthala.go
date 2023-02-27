package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达摩萨他罗DharmasthalaBarony struct {
	feud.BaseBarony
}

var BDharmasthala达摩萨他罗 feud.Barony = &达摩萨他罗DharmasthalaBarony{}

func init() {
    f := BDharmasthala达摩萨他罗.(*达摩萨他罗DharmasthalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharmasthala",
		TitleName: "达摩萨他罗",
		TitleCode: "b_dharmasthala",
	}
}
