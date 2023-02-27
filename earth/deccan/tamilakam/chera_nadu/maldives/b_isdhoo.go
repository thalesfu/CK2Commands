package maldives

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯杜IsdhooBarony struct {
	feud.BaseBarony
}

var BIsdhoo伊斯杜 feud.Barony = &伊斯杜IsdhooBarony{}

func init() {
    f := BIsdhoo伊斯杜.(*伊斯杜IsdhooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "isdhoo",
		TitleName: "伊斯杜",
		TitleCode: "b_isdhoo",
	}
}
