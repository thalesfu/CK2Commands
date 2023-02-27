package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌芬海姆UffenheimBarony struct {
	feud.BaseBarony
}

var BUffenheim乌芬海姆 feud.Barony = &乌芬海姆UffenheimBarony{}

func init() {
    f := BUffenheim乌芬海姆.(*乌芬海姆UffenheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uffenheim",
		TitleName: "乌芬海姆",
		TitleCode: "b_uffenheim",
	}
}
