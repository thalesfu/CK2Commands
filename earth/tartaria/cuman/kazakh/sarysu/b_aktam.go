package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克塔姆AktamBarony struct {
	feud.BaseBarony
}

var BAktam阿克塔姆 feud.Barony = &阿克塔姆AktamBarony{}

func init() {
    f := BAktam阿克塔姆.(*阿克塔姆AktamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aktam",
		TitleName: "阿克塔姆",
		TitleCode: "b_aktam",
	}
}
