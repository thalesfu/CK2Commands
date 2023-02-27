package krain

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣法伊特StveitBarony struct {
	feud.BaseBarony
}

var BStveit圣法伊特 feud.Barony = &圣法伊特StveitBarony{}

func init() {
    f := BStveit圣法伊特.(*圣法伊特StveitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stveit",
		TitleName: "圣法伊特",
		TitleCode: "b_stveit",
	}
}
