package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌拉拉UlalaBarony struct {
	feud.BaseBarony
}

var BUlala乌拉拉 feud.Barony = &乌拉拉UlalaBarony{}

func init() {
	f := BUlala乌拉拉.(*乌拉拉UlalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulala",
		TitleName: "乌拉拉",
		TitleCode: "b_ulala",
	}
}
