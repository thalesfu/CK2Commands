package uch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 招陀利ChaudariBarony struct {
	feud.BaseBarony
}

var BChaudari招陀利 feud.Barony = &招陀利ChaudariBarony{}

func init() {
    f := BChaudari招陀利.(*招陀利ChaudariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaudari",
		TitleName: "招陀利",
		TitleCode: "b_chaudari",
	}
}
