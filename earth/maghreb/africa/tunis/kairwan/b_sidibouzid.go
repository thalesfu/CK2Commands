package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪布济德SidibouzidBarony struct {
	feud.BaseBarony
}

var BSidibouzid西迪布济德 feud.Barony = &西迪布济德SidibouzidBarony{}

func init() {
	f := BSidibouzid西迪布济德.(*西迪布济德SidibouzidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidibouzid",
		TitleName: "西迪布济德",
		TitleCode: "b_sidibouzid",
	}
}
