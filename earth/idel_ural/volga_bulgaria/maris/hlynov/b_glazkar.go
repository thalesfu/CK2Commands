package hlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉兹卡尔GlazkarBarony struct {
	feud.BaseBarony
}

var BGlazkar格拉兹卡尔 feud.Barony = &格拉兹卡尔GlazkarBarony{}

func init() {
    f := BGlazkar格拉兹卡尔.(*格拉兹卡尔GlazkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glazkar",
		TitleName: "格拉兹卡尔",
		TitleCode: "b_glazkar",
	}
}
