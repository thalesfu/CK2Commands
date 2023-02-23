package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞代SedehBarony struct {
	feud.BaseBarony
}

var BSedeh塞代 feud.Barony = &塞代SedehBarony{}

func init() {
	f := BSedeh塞代.(*塞代SedehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sedeh",
		TitleName: "塞代",
		TitleCode: "b_sedeh",
	}
}
