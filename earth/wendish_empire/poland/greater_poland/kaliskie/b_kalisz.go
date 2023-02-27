package kaliskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利什KaliszBarony struct {
	feud.BaseBarony
}

var BKalisz卡利什 feud.Barony = &卡利什KaliszBarony{}

func init() {
    f := BKalisz卡利什.(*卡利什KaliszBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalisz",
		TitleName: "卡利什",
		TitleCode: "b_kalisz",
	}
}
