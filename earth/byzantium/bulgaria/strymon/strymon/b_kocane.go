package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科查内KocaneBarony struct {
	feud.BaseBarony
}

var BKocane科查内 feud.Barony = &科查内KocaneBarony{}

func init() {
    f := BKocane科查内.(*科查内KocaneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kocane",
		TitleName: "科查内",
		TitleCode: "b_kocane",
	}
}
