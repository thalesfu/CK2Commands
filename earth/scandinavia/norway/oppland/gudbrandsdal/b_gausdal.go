package gudbrandsdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖于斯达尔GausdalBarony struct {
	feud.BaseBarony
}

var BGausdal盖于斯达尔 feud.Barony = &盖于斯达尔GausdalBarony{}

func init() {
    f := BGausdal盖于斯达尔.(*盖于斯达尔GausdalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gausdal",
		TitleName: "盖于斯达尔",
		TitleCode: "b_gausdal",
	}
}
