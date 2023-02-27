package clydesdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉纳克LanarkBarony struct {
	feud.BaseBarony
}

var BLanark拉纳克 feud.Barony = &拉纳克LanarkBarony{}

func init() {
    f := BLanark拉纳克.(*拉纳克LanarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lanark",
		TitleName: "拉纳克",
		TitleCode: "b_lanark",
	}
}
