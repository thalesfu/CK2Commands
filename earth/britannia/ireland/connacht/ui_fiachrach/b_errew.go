package ui_fiachrach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃拉ErrewBarony struct {
	feud.BaseBarony
}

var BErrew埃拉 feud.Barony = &埃拉ErrewBarony{}

func init() {
    f := BErrew埃拉.(*埃拉ErrewBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "errew",
		TitleName: "埃拉",
		TitleCode: "b_errew",
	}
}
