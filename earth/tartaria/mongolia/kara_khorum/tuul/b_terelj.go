package tuul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特勒尔吉TereljBarony struct {
	feud.BaseBarony
}

var BTerelj特勒尔吉 feud.Barony = &特勒尔吉TereljBarony{}

func init() {
    f := BTerelj特勒尔吉.(*特勒尔吉TereljBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terelj",
		TitleName: "特勒尔吉",
		TitleCode: "b_terelj",
	}
}
