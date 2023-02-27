package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 套多斯TaoudouchBarony struct {
	feud.BaseBarony
}

var BTaoudouch套多斯 feud.Barony = &套多斯TaoudouchBarony{}

func init() {
    f := BTaoudouch套多斯.(*套多斯TaoudouchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taoudouch",
		TitleName: "套多斯",
		TitleCode: "b_taoudouch",
	}
}
