package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙什德SasdBarony struct {
	feud.BaseBarony
}

var BSasd沙什德 feud.Barony = &沙什德SasdBarony{}

func init() {
    f := BSasd沙什德.(*沙什德SasdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sasd",
		TitleName: "沙什德",
		TitleCode: "b_sasd",
	}
}
