package aland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃克勒EckeroBarony struct {
	feud.BaseBarony
}

var BEckero埃克勒 feud.Barony = &埃克勒EckeroBarony{}

func init() {
	f := BEckero埃克勒.(*埃克勒EckeroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eckero",
		TitleName: "埃克勒",
		TitleCode: "b_eckero",
	}
}
