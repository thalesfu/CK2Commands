package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉茨拉夫BratslavBarony struct {
	feud.BaseBarony
}

var BBratslav布拉茨拉夫 feud.Barony = &布拉茨拉夫BratslavBarony{}

func init() {
    f := BBratslav布拉茨拉夫.(*布拉茨拉夫BratslavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bratslav",
		TitleName: "布拉茨拉夫",
		TitleCode: "b_bratslav",
	}
}
