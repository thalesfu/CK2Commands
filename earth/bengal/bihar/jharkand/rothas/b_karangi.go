package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽蓝耆KarangiBarony struct {
	feud.BaseBarony
}

var BKarangi伽蓝耆 feud.Barony = &伽蓝耆KarangiBarony{}

func init() {
    f := BKarangi伽蓝耆.(*伽蓝耆KarangiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karangi",
		TitleName: "伽蓝耆",
		TitleCode: "b_karangi",
	}
}
