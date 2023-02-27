package boqtybay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古别尔林斯基GuberlinskiyBarony struct {
	feud.BaseBarony
}

var BGuberlinskiy古别尔林斯基 feud.Barony = &古别尔林斯基GuberlinskiyBarony{}

func init() {
    f := BGuberlinskiy古别尔林斯基.(*古别尔林斯基GuberlinskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guberlinskiy",
		TitleName: "古别尔林斯基",
		TitleCode: "b_guberlinskiy",
	}
}
