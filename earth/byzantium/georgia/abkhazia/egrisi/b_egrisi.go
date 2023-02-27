package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃格里西EgrisiBarony struct {
	feud.BaseBarony
}

var BEgrisi埃格里西 feud.Barony = &埃格里西EgrisiBarony{}

func init() {
    f := BEgrisi埃格里西.(*埃格里西EgrisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "egrisi",
		TitleName: "埃格里西",
		TitleCode: "b_egrisi",
	}
}
