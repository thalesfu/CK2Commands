package chlisselbourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛杰伊诺耶波列Lodeynoye_poleBarony struct {
	feud.BaseBarony
}

var BLodeynoye_pole洛杰伊诺耶波列 feud.Barony = &洛杰伊诺耶波列Lodeynoye_poleBarony{}

func init() {
    f := BLodeynoye_pole洛杰伊诺耶波列.(*洛杰伊诺耶波列Lodeynoye_poleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lodeynoye_pole",
		TitleName: "洛杰伊诺耶波列",
		TitleCode: "b_lodeynoye_pole",
	}
}
