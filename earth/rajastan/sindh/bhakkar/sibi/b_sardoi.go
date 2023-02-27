package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨豆伊SardoiBarony struct {
	feud.BaseBarony
}

var BSardoi萨豆伊 feud.Barony = &萨豆伊SardoiBarony{}

func init() {
    f := BSardoi萨豆伊.(*萨豆伊SardoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sardoi",
		TitleName: "萨豆伊",
		TitleCode: "b_sardoi",
	}
}
