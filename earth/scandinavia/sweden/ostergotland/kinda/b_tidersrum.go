package kinda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂德什鲁姆TidersrumBarony struct {
	feud.BaseBarony
}

var BTidersrum蒂德什鲁姆 feud.Barony = &蒂德什鲁姆TidersrumBarony{}

func init() {
	f := BTidersrum蒂德什鲁姆.(*蒂德什鲁姆TidersrumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tidersrum",
		TitleName: "蒂德什鲁姆",
		TitleCode: "b_tidersrum",
	}
}
