package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇斯托耶波列Chistoye_poleBarony struct {
	feud.BaseBarony
}

var BChistoye_pole奇斯托耶波列 feud.Barony = &奇斯托耶波列Chistoye_poleBarony{}

func init() {
    f := BChistoye_pole奇斯托耶波列.(*奇斯托耶波列Chistoye_poleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chistoye_pole",
		TitleName: "奇斯托耶波列",
		TitleCode: "b_chistoye_pole",
	}
}
