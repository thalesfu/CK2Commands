package magdeburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝恩堡BernburgBarony struct {
	feud.BaseBarony
}

var BBernburg贝恩堡 feud.Barony = &贝恩堡BernburgBarony{}

func init() {
    f := BBernburg贝恩堡.(*贝恩堡BernburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bernburg",
		TitleName: "贝恩堡",
		TitleCode: "b_bernburg",
	}
}
