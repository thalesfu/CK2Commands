package balkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾哈努姆AlkhanoumBarony struct {
	feud.BaseBarony
}

var BAlkhanoum艾哈努姆 feud.Barony = &艾哈努姆AlkhanoumBarony{}

func init() {
	f := BAlkhanoum艾哈努姆.(*艾哈努姆AlkhanoumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkhanoum",
		TitleName: "艾哈努姆",
		TitleCode: "b_alkhanoum",
	}
}
