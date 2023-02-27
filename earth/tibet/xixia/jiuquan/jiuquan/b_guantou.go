package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 观头GuantouBarony struct {
	feud.BaseBarony
}

var BGuantou观头 feud.Barony = &观头GuantouBarony{}

func init() {
    f := BGuantou观头.(*观头GuantouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guantou",
		TitleName: "观头",
		TitleCode: "b_guantou",
	}
}
