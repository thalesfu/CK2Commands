package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索斯尼齐亚SosnytsiaBarony struct {
	feud.BaseBarony
}

var BSosnytsia索斯尼齐亚 feud.Barony = &索斯尼齐亚SosnytsiaBarony{}

func init() {
    f := BSosnytsia索斯尼齐亚.(*索斯尼齐亚SosnytsiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sosnytsia",
		TitleName: "索斯尼齐亚",
		TitleCode: "b_sosnytsia",
	}
}
