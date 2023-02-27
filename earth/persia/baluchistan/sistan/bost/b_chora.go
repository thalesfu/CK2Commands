package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔拉ChoraBarony struct {
	feud.BaseBarony
}

var BChora乔拉 feud.Barony = &乔拉ChoraBarony{}

func init() {
    f := BChora乔拉.(*乔拉ChoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chora",
		TitleName: "乔拉",
		TitleCode: "b_chora",
	}
}
