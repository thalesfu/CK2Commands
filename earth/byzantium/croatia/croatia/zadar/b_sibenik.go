package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希贝尼克SibenikBarony struct {
	feud.BaseBarony
}

var BSibenik希贝尼克 feud.Barony = &希贝尼克SibenikBarony{}

func init() {
	f := BSibenik希贝尼克.(*希贝尼克SibenikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sibenik",
		TitleName: "希贝尼克",
		TitleCode: "b_sibenik",
	}
}
