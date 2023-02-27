package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 撒马尔罕SamarkandBarony struct {
	feud.BaseBarony
}

var BSamarkand撒马尔罕 feud.Barony = &撒马尔罕SamarkandBarony{}

func init() {
    f := BSamarkand撒马尔罕.(*撒马尔罕SamarkandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samarkand",
		TitleName: "撒马尔罕",
		TitleCode: "b_samarkand",
	}
}
