package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 派尔努ParnawBarony struct {
	feud.BaseBarony
}

var BParnaw派尔努 feud.Barony = &派尔努ParnawBarony{}

func init() {
	f := BParnaw派尔努.(*派尔努ParnawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parnaw",
		TitleName: "派尔努",
		TitleCode: "b_parnaw",
	}
}
