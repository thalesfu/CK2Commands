package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌特里利亚斯UtrillasBarony struct {
	feud.BaseBarony
}

var BUtrillas乌特里利亚斯 feud.Barony = &乌特里利亚斯UtrillasBarony{}

func init() {
    f := BUtrillas乌特里利亚斯.(*乌特里利亚斯UtrillasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "utrillas",
		TitleName: "乌特里利亚斯",
		TitleCode: "b_utrillas",
	}
}
