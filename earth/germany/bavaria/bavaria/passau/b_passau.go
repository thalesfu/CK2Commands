package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕绍PassauBarony struct {
	feud.BaseBarony
}

var BPassau帕绍 feud.Barony = &帕绍PassauBarony{}

func init() {
    f := BPassau帕绍.(*帕绍PassauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "passau",
		TitleName: "帕绍",
		TitleCode: "b_passau",
	}
}
