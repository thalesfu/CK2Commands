package starodub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维申基VyshenkyBarony struct {
	feud.BaseBarony
}

var BVyshenky维申基 feud.Barony = &维申基VyshenkyBarony{}

func init() {
    f := BVyshenky维申基.(*维申基VyshenkyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyshenky",
		TitleName: "维申基",
		TitleCode: "b_vyshenky",
	}
}
