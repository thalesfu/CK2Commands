package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏克舍赫克SukelsheyuhkBarony struct {
	feud.BaseBarony
}

var BSukelsheyuhk苏克舍赫克 feud.Barony = &苏克舍赫克SukelsheyuhkBarony{}

func init() {
	f := BSukelsheyuhk苏克舍赫克.(*苏克舍赫克SukelsheyuhkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sukelsheyuhk",
		TitleName: "苏克舍赫克",
		TitleCode: "b_sukelsheyuhk",
	}
}
