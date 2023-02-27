package keltma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克尔乔姆亚KerchomyaBarony struct {
	feud.BaseBarony
}

var BKerchomya克尔乔姆亚 feud.Barony = &克尔乔姆亚KerchomyaBarony{}

func init() {
    f := BKerchomya克尔乔姆亚.(*克尔乔姆亚KerchomyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kerchomya",
		TitleName: "克尔乔姆亚",
		TitleCode: "b_kerchomya",
	}
}
