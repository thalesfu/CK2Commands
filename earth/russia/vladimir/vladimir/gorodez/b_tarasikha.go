package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉西哈TarasikhaBarony struct {
	feud.BaseBarony
}

var BTarasikha塔拉西哈 feud.Barony = &塔拉西哈TarasikhaBarony{}

func init() {
    f := BTarasikha塔拉西哈.(*塔拉西哈TarasikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarasikha",
		TitleName: "塔拉西哈",
		TitleCode: "b_tarasikha",
	}
}
