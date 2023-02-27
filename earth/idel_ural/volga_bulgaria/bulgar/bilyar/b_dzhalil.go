package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾利利DzhalilBarony struct {
	feud.BaseBarony
}

var BDzhalil贾利利 feud.Barony = &贾利利DzhalilBarony{}

func init() {
    f := BDzhalil贾利利.(*贾利利DzhalilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dzhalil",
		TitleName: "贾利利",
		TitleCode: "b_dzhalil",
	}
}
