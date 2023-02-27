package uda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 结雅斯基ZeyskyBarony struct {
	feud.BaseBarony
}

var BZeysky结雅斯基 feud.Barony = &结雅斯基ZeyskyBarony{}

func init() {
    f := BZeysky结雅斯基.(*结雅斯基ZeyskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeysky",
		TitleName: "结雅斯基",
		TitleCode: "b_zeysky",
	}
}
