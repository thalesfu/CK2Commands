package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 本哈穆什BenhammouchBarony struct {
	feud.BaseBarony
}

var BBenhammouch本哈穆什 feud.Barony = &本哈穆什BenhammouchBarony{}

func init() {
	f := BBenhammouch本哈穆什.(*本哈穆什BenhammouchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benhammouch",
		TitleName: "本哈穆什",
		TitleCode: "b_benhammouch",
	}
}
