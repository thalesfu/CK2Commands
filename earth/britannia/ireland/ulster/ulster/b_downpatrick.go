package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 唐帕特里克DownpatrickBarony struct {
	feud.BaseBarony
}

var BDownpatrick唐帕特里克 feud.Barony = &唐帕特里克DownpatrickBarony{}

func init() {
    f := BDownpatrick唐帕特里克.(*唐帕特里克DownpatrickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "downpatrick",
		TitleName: "唐帕特里克",
		TitleCode: "b_downpatrick",
	}
}
