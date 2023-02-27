package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩季_亚赫PytyakhBarony struct {
	feud.BaseBarony
}

var BPytyakh佩季_亚赫 feud.Barony = &佩季_亚赫PytyakhBarony{}

func init() {
    f := BPytyakh佩季_亚赫.(*佩季_亚赫PytyakhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pytyakh",
		TitleName: "佩季_亚赫",
		TitleCode: "b_pytyakh",
	}
}
