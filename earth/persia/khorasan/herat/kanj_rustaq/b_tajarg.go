package kanj_rustaq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔贾格TajargBarony struct {
	feud.BaseBarony
}

var BTajarg塔贾格 feud.Barony = &塔贾格TajargBarony{}

func init() {
    f := BTajarg塔贾格.(*塔贾格TajargBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tajarg",
		TitleName: "塔贾格",
		TitleCode: "b_tajarg",
	}
}
