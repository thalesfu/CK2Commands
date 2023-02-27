package bira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉拉TellaBarony struct {
	feud.BaseBarony
}

var BTella特拉拉 feud.Barony = &特拉拉TellaBarony{}

func init() {
    f := BTella特拉拉.(*特拉拉TellaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tella",
		TitleName: "特拉拉",
		TitleCode: "b_tella",
	}
}
