package izhma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊日马IzhmaBarony struct {
	feud.BaseBarony
}

var BIzhma伊日马 feud.Barony = &伊日马IzhmaBarony{}

func init() {
    f := BIzhma伊日马.(*伊日马IzhmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "izhma",
		TitleName: "伊日马",
		TitleCode: "b_izhma",
	}
}
