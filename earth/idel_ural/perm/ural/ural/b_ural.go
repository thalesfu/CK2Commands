package ural

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌拉尔UralBarony struct {
	feud.BaseBarony
}

var BUral乌拉尔 feud.Barony = &乌拉尔UralBarony{}

func init() {
    f := BUral乌拉尔.(*乌拉尔UralBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ural",
		TitleName: "乌拉尔",
		TitleCode: "b_ural",
	}
}
