package trigarta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇烂达罗JalandharBarony struct {
	feud.BaseBarony
}

var BJalandhar阇烂达罗 feud.Barony = &阇烂达罗JalandharBarony{}

func init() {
	f := BJalandhar阇烂达罗.(*阇烂达罗JalandharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jalandhar",
		TitleName: "阇烂达罗",
		TitleCode: "b_jalandhar",
	}
}
