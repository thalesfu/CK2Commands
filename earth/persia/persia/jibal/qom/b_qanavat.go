package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀纳瓦特QanavatBarony struct {
	feud.BaseBarony
}

var BQanavat喀纳瓦特 feud.Barony = &喀纳瓦特QanavatBarony{}

func init() {
	f := BQanavat喀纳瓦特.(*喀纳瓦特QanavatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qanavat",
		TitleName: "喀纳瓦特",
		TitleCode: "b_qanavat",
	}
}
