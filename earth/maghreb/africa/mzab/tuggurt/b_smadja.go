package tuggurt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯马贾SmadjaBarony struct {
	feud.BaseBarony
}

var BSmadja斯马贾 feud.Barony = &斯马贾SmadjaBarony{}

func init() {
	f := BSmadja斯马贾.(*斯马贾SmadjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smadja",
		TitleName: "斯马贾",
		TitleCode: "b_smadja",
	}
}
