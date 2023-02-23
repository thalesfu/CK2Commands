package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季耶UsteBarony struct {
	feud.BaseBarony
}

var BUste乌斯季耶 feud.Barony = &乌斯季耶UsteBarony{}

func init() {
	f := BUste乌斯季耶.(*乌斯季耶UsteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uste",
		TitleName: "乌斯季耶",
		TitleCode: "b_uste",
	}
}
