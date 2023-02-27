package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托拉ToraBarony struct {
	feud.BaseBarony
}

var BTora托拉 feud.Barony = &托拉ToraBarony{}

func init() {
    f := BTora托拉.(*托拉ToraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tora",
		TitleName: "托拉",
		TitleCode: "b_tora",
	}
}
