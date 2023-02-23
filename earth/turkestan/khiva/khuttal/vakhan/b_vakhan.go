package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 镬侃VakhanBarony struct {
	feud.BaseBarony
}

var BVakhan镬侃 feud.Barony = &镬侃VakhanBarony{}

func init() {
	f := BVakhan镬侃.(*镬侃VakhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vakhan",
		TitleName: "镬侃",
		TitleCode: "b_vakhan",
	}
}
