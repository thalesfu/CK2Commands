package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌格利奇波列Ugliche_poleBarony struct {
	feud.BaseBarony
}

var BUgliche_pole乌格利奇波列 feud.Barony = &乌格利奇波列Ugliche_poleBarony{}

func init() {
    f := BUgliche_pole乌格利奇波列.(*乌格利奇波列Ugliche_poleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ugliche_pole",
		TitleName: "乌格利奇波列",
		TitleCode: "b_ugliche_pole",
	}
}
