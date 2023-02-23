package claudiopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑加里奥斯SangariusBarony struct {
	feud.BaseBarony
}

var BSangarius桑加里奥斯 feud.Barony = &桑加里奥斯SangariusBarony{}

func init() {
	f := BSangarius桑加里奥斯.(*桑加里奥斯SangariusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangarius",
		TitleName: "桑加里奥斯",
		TitleCode: "b_sangarius",
	}
}
