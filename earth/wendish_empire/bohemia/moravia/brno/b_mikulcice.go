package brno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米库尔奇采MikulciceBarony struct {
	feud.BaseBarony
}

var BMikulcice米库尔奇采 feud.Barony = &米库尔奇采MikulciceBarony{}

func init() {
    f := BMikulcice米库尔奇采.(*米库尔奇采MikulciceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mikulcice",
		TitleName: "米库尔奇采",
		TitleCode: "b_mikulcice",
	}
}
