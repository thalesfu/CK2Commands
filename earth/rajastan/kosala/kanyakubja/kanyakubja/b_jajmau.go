package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇祇牟JajmauBarony struct {
	feud.BaseBarony
}

var BJajmau阇祇牟 feud.Barony = &阇祇牟JajmauBarony{}

func init() {
	f := BJajmau阇祇牟.(*阇祇牟JajmauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jajmau",
		TitleName: "阇祇牟",
		TitleCode: "b_jajmau",
	}
}
