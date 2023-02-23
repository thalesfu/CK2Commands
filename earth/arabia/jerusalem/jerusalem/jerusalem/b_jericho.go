package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰里科JerichoBarony struct {
	feud.BaseBarony
}

var BJericho杰里科 feud.Barony = &杰里科JerichoBarony{}

func init() {
	f := BJericho杰里科.(*杰里科JerichoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jericho",
		TitleName: "杰里科",
		TitleCode: "b_jericho",
	}
}
