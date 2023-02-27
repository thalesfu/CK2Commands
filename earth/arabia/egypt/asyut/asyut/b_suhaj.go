package asyut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏哈吉SuhajBarony struct {
	feud.BaseBarony
}

var BSuhaj苏哈吉 feud.Barony = &苏哈吉SuhajBarony{}

func init() {
    f := BSuhaj苏哈吉.(*苏哈吉SuhajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suhaj",
		TitleName: "苏哈吉",
		TitleCode: "b_suhaj",
	}
}
