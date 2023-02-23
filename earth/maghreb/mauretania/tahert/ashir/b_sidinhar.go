package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪恩哈尔SidinharBarony struct {
	feud.BaseBarony
}

var BSidinhar西迪恩哈尔 feud.Barony = &西迪恩哈尔SidinharBarony{}

func init() {
	f := BSidinhar西迪恩哈尔.(*西迪恩哈尔SidinharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidinhar",
		TitleName: "西迪恩哈尔",
		TitleCode: "b_sidinhar",
	}
}
