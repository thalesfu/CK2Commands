package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀内提DhanetiBarony struct {
	feud.BaseBarony
}

var BDhaneti陀内提 feud.Barony = &陀内提DhanetiBarony{}

func init() {
    f := BDhaneti陀内提.(*陀内提DhanetiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhaneti",
		TitleName: "陀内提",
		TitleCode: "b_dhaneti",
	}
}
