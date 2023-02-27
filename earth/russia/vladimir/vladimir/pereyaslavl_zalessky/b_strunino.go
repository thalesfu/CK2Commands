package pereyaslavl_zalessky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特鲁尼诺StruninoBarony struct {
	feud.BaseBarony
}

var BStrunino斯特鲁尼诺 feud.Barony = &斯特鲁尼诺StruninoBarony{}

func init() {
    f := BStrunino斯特鲁尼诺.(*斯特鲁尼诺StruninoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "strunino",
		TitleName: "斯特鲁尼诺",
		TitleCode: "b_strunino",
	}
}
