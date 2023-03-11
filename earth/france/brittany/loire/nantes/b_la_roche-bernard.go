package nantes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉罗什_贝尔纳La_roche_bernardBarony struct {
	feud.BaseBarony
}

var BLa_roche_bernard拉罗什_贝尔纳 feud.Barony = &拉罗什_贝尔纳La_roche_bernardBarony{}

func init() {
    f := BLa_roche_bernard拉罗什_贝尔纳.(*拉罗什_贝尔纳La_roche_bernardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "la_roche_bernard",
		TitleName: "拉罗什_贝尔纳",
		TitleCode: "b_la_roche-bernard",
	}
}
