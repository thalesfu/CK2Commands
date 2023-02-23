package vijnot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利兜阿SaitualBarony struct {
	feud.BaseBarony
}

var BSaitual室利兜阿 feud.Barony = &室利兜阿SaitualBarony{}

func init() {
	f := BSaitual室利兜阿.(*室利兜阿SaitualBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saitual",
		TitleName: "室利兜阿",
		TitleCode: "b_saitual",
	}
}
