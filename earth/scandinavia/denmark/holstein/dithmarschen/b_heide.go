package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海德HeideBarony struct {
	feud.BaseBarony
}

var BHeide海德 feud.Barony = &海德HeideBarony{}

func init() {
	f := BHeide海德.(*海德HeideBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "heide",
		TitleName: "海德",
		TitleCode: "b_heide",
	}
}
