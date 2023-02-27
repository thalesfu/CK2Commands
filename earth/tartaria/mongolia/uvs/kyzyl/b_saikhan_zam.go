package kyzyl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赛罕扎木Saikhan_zamBarony struct {
	feud.BaseBarony
}

var BSaikhan_zam赛罕扎木 feud.Barony = &赛罕扎木Saikhan_zamBarony{}

func init() {
    f := BSaikhan_zam赛罕扎木.(*赛罕扎木Saikhan_zamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saikhan_zam",
		TitleName: "赛罕扎木",
		TitleCode: "b_saikhan_zam",
	}
}
