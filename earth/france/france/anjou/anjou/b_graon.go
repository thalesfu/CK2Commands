package anjou

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格朗GraonBarony struct {
	feud.BaseBarony
}

var BGraon格朗 feud.Barony = &格朗GraonBarony{}

func init() {
	f := BGraon格朗.(*格朗GraonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "graon",
		TitleName: "格朗",
		TitleCode: "b_graon",
	}
}
