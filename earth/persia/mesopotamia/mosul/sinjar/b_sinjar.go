package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛贾尔SinjarBarony struct {
	feud.BaseBarony
}

var BSinjar辛贾尔 feud.Barony = &辛贾尔SinjarBarony{}

func init() {
	f := BSinjar辛贾尔.(*辛贾尔SinjarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinjar",
		TitleName: "辛贾尔",
		TitleCode: "b_sinjar",
	}
}
