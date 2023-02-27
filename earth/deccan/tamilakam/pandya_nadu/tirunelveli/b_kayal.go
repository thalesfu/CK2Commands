package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽一KayalBarony struct {
	feud.BaseBarony
}

var BKayal伽一 feud.Barony = &伽一KayalBarony{}

func init() {
    f := BKayal伽一.(*伽一KayalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kayal",
		TitleName: "伽一",
		TitleCode: "b_kayal",
	}
}
