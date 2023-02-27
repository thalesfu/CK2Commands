package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴巴尤尔特BabayurtBarony struct {
	feud.BaseBarony
}

var BBabayurt巴巴尤尔特 feud.Barony = &巴巴尤尔特BabayurtBarony{}

func init() {
    f := BBabayurt巴巴尤尔特.(*巴巴尤尔特BabayurtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babayurt",
		TitleName: "巴巴尤尔特",
		TitleCode: "b_babayurt",
	}
}
