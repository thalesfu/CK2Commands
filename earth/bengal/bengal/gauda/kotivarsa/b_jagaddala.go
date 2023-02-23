package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇伽陀罗JagaddalaBarony struct {
	feud.BaseBarony
}

var BJagaddala阇伽陀罗 feud.Barony = &阇伽陀罗JagaddalaBarony{}

func init() {
	f := BJagaddala阇伽陀罗.(*阇伽陀罗JagaddalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jagaddala",
		TitleName: "阇伽陀罗",
		TitleCode: "b_jagaddala",
	}
}
