package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古姆萨GumsaBarony struct {
	feud.BaseBarony
}

var BGumsa古姆萨 feud.Barony = &古姆萨GumsaBarony{}

func init() {
	f := BGumsa古姆萨.(*古姆萨GumsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gumsa",
		TitleName: "古姆萨",
		TitleCode: "b_gumsa",
	}
}
