package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢基夫LukivBarony struct {
	feud.BaseBarony
}

var BLukiv卢基夫 feud.Barony = &卢基夫LukivBarony{}

func init() {
    f := BLukiv卢基夫.(*卢基夫LukivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lukiv",
		TitleName: "卢基夫",
		TitleCode: "b_lukiv",
	}
}
