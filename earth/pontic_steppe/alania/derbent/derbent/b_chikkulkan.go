package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇库勒坎ChikkulkanBarony struct {
	feud.BaseBarony
}

var BChikkulkan奇库勒坎 feud.Barony = &奇库勒坎ChikkulkanBarony{}

func init() {
    f := BChikkulkan奇库勒坎.(*奇库勒坎ChikkulkanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chikkulkan",
		TitleName: "奇库勒坎",
		TitleCode: "b_chikkulkan",
	}
}
