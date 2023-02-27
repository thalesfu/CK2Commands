package bacs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲍尧BajaBarony struct {
	feud.BaseBarony
}

var BBaja鲍尧 feud.Barony = &鲍尧BajaBarony{}

func init() {
    f := BBaja鲍尧.(*鲍尧BajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baja",
		TitleName: "鲍尧",
		TitleCode: "b_baja",
	}
}
