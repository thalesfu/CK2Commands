package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德汝DerokBarony struct {
	feud.BaseBarony
}

var BDerok德汝 feud.Barony = &德汝DerokBarony{}

func init() {
	f := BDerok德汝.(*德汝DerokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "derok",
		TitleName: "德汝",
		TitleCode: "b_derok",
	}
}
