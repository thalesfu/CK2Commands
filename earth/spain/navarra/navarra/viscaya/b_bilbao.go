package viscaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毕尔巴鄂BilbaoBarony struct {
	feud.BaseBarony
}

var BBilbao毕尔巴鄂 feud.Barony = &毕尔巴鄂BilbaoBarony{}

func init() {
	f := BBilbao毕尔巴鄂.(*毕尔巴鄂BilbaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilbao",
		TitleName: "毕尔巴鄂",
		TitleCode: "b_bilbao",
	}
}
