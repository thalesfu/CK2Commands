package vesoul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博库尔BeaucourtBarony struct {
	feud.BaseBarony
}

var BBeaucourt博库尔 feud.Barony = &博库尔BeaucourtBarony{}

func init() {
    f := BBeaucourt博库尔.(*博库尔BeaucourtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beaucourt",
		TitleName: "博库尔",
		TitleCode: "b_beaucourt",
	}
}
