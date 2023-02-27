package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科隆比耶ColombierBarony struct {
	feud.BaseBarony
}

var BColombier科隆比耶 feud.Barony = &科隆比耶ColombierBarony{}

func init() {
    f := BColombier科隆比耶.(*科隆比耶ColombierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "colombier",
		TitleName: "科隆比耶",
		TitleCode: "b_colombier",
	}
}
