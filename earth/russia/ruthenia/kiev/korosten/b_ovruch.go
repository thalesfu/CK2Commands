package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥夫鲁奇OvruchBarony struct {
	feud.BaseBarony
}

var BOvruch奥夫鲁奇 feud.Barony = &奥夫鲁奇OvruchBarony{}

func init() {
	f := BOvruch奥夫鲁奇.(*奥夫鲁奇OvruchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ovruch",
		TitleName: "奥夫鲁奇",
		TitleCode: "b_ovruch",
	}
}
