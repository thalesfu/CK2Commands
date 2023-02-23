package mahdia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪曼苏尔SidimansourBarony struct {
	feud.BaseBarony
}

var BSidimansour西迪曼苏尔 feud.Barony = &西迪曼苏尔SidimansourBarony{}

func init() {
	f := BSidimansour西迪曼苏尔.(*西迪曼苏尔SidimansourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidimansour",
		TitleName: "西迪曼苏尔",
		TitleCode: "b_sidimansour",
	}
}
