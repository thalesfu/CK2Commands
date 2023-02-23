package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉多LhatokBarony struct {
	feud.BaseBarony
}

var BLhatok拉多 feud.Barony = &拉多LhatokBarony{}

func init() {
	f := BLhatok拉多.(*拉多LhatokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lhatok",
		TitleName: "拉多",
		TitleCode: "b_lhatok",
	}
}
