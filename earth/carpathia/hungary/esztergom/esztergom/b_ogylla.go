package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧焦洛OgyllaBarony struct {
	feud.BaseBarony
}

var BOgylla欧焦洛 feud.Barony = &欧焦洛OgyllaBarony{}

func init() {
	f := BOgylla欧焦洛.(*欧焦洛OgyllaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ogylla",
		TitleName: "欧焦洛",
		TitleCode: "b_ogylla",
	}
}
