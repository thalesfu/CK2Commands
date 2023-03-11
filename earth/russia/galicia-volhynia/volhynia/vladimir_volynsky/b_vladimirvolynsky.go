package vladimir_volynsky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉基米尔沃伦斯基VladimirvolynskyBarony struct {
	feud.BaseBarony
}

var BVladimirvolynsky弗拉基米尔沃伦斯基 feud.Barony = &弗拉基米尔沃伦斯基VladimirvolynskyBarony{}

func init() {
    f := BVladimirvolynsky弗拉基米尔沃伦斯基.(*弗拉基米尔沃伦斯基VladimirvolynskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vladimirvolynsky",
		TitleName: "弗拉基米尔沃伦斯基",
		TitleCode: "b_vladimirvolynsky",
	}
}
