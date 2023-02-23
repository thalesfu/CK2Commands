package furstenberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多瑙埃兴根DonaueschingenBarony struct {
	feud.BaseBarony
}

var BDonaueschingen多瑙埃兴根 feud.Barony = &多瑙埃兴根DonaueschingenBarony{}

func init() {
	f := BDonaueschingen多瑙埃兴根.(*多瑙埃兴根DonaueschingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donaueschingen",
		TitleName: "多瑙埃兴根",
		TitleCode: "b_donaueschingen",
	}
}
