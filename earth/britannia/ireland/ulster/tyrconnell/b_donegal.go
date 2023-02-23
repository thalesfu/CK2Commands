package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多尼戈尔DonegalBarony struct {
	feud.BaseBarony
}

var BDonegal多尼戈尔 feud.Barony = &多尼戈尔DonegalBarony{}

func init() {
	f := BDonegal多尼戈尔.(*多尼戈尔DonegalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donegal",
		TitleName: "多尼戈尔",
		TitleCode: "b_donegal",
	}
}
