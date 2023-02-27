package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 包岑BautzenBarony struct {
	feud.BaseBarony
}

var BBautzen包岑 feud.Barony = &包岑BautzenBarony{}

func init() {
    f := BBautzen包岑.(*包岑BautzenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bautzen",
		TitleName: "包岑",
		TitleCode: "b_bautzen",
	}
}
