package yamalia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊图_亚哈ItuyakhaBarony struct {
	feud.BaseBarony
}

var BItuyakha伊图_亚哈 feud.Barony = &伊图_亚哈ItuyakhaBarony{}

func init() {
	f := BItuyakha伊图_亚哈.(*伊图_亚哈ItuyakhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ituyakha",
		TitleName: "伊图_亚哈",
		TitleCode: "b_ituyakha",
	}
}
