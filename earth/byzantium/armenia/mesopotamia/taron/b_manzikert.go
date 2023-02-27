package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼齐刻尔特ManzikertBarony struct {
	feud.BaseBarony
}

var BManzikert曼齐刻尔特 feud.Barony = &曼齐刻尔特ManzikertBarony{}

func init() {
    f := BManzikert曼齐刻尔特.(*曼齐刻尔特ManzikertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manzikert",
		TitleName: "曼齐刻尔特",
		TitleCode: "b_manzikert",
	}
}
