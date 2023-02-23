package rashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹韦钱ZvecanBarony struct {
	feud.BaseBarony
}

var BZvecan兹韦钱 feud.Barony = &兹韦钱ZvecanBarony{}

func init() {
	f := BZvecan兹韦钱.(*兹韦钱ZvecanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zvecan",
		TitleName: "兹韦钱",
		TitleCode: "b_zvecan",
	}
}
