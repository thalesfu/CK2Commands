package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特罗法迪亚StrofadiaBarony struct {
	feud.BaseBarony
}

var BStrofadia斯特罗法迪亚 feud.Barony = &斯特罗法迪亚StrofadiaBarony{}

func init() {
    f := BStrofadia斯特罗法迪亚.(*斯特罗法迪亚StrofadiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strofadia",
		TitleName: "斯特罗法迪亚",
		TitleCode: "b_strofadia",
	}
}
