package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇羯罗JakarBarony struct {
	feud.BaseBarony
}

var BJakar阇羯罗 feud.Barony = &阇羯罗JakarBarony{}

func init() {
	f := BJakar阇羯罗.(*阇羯罗JakarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jakar",
		TitleName: "阇羯罗",
		TitleCode: "b_jakar",
	}
}
