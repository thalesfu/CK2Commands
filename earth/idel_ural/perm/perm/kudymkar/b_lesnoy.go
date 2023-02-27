package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列斯诺伊LesnoyBarony struct {
	feud.BaseBarony
}

var BLesnoy列斯诺伊 feud.Barony = &列斯诺伊LesnoyBarony{}

func init() {
    f := BLesnoy列斯诺伊.(*列斯诺伊LesnoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lesnoy",
		TitleName: "列斯诺伊",
		TitleCode: "b_lesnoy",
	}
}
