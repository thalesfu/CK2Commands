package fitri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗昂TroanBarony struct {
	feud.BaseBarony
}

var BTroan特罗昂 feud.Barony = &特罗昂TroanBarony{}

func init() {
    f := BTroan特罗昂.(*特罗昂TroanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "troan",
		TitleName: "特罗昂",
		TitleCode: "b_troan",
	}
}
