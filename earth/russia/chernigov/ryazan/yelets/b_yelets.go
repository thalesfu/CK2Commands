package yelets

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶列茨YeletsBarony struct {
	feud.BaseBarony
}

var BYelets叶列茨 feud.Barony = &叶列茨YeletsBarony{}

func init() {
    f := BYelets叶列茨.(*叶列茨YeletsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yelets",
		TitleName: "叶列茨",
		TitleCode: "b_yelets",
	}
}
