package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙普ShapBarony struct {
	feud.BaseBarony
}

var BShap沙普 feud.Barony = &沙普ShapBarony{}

func init() {
	f := BShap沙普.(*沙普ShapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shap",
		TitleName: "沙普",
		TitleCode: "b_shap",
	}
}
