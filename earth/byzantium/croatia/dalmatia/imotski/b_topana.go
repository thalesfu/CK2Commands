package imotski

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托帕纳TopanaBarony struct {
	feud.BaseBarony
}

var BTopana托帕纳 feud.Barony = &托帕纳TopanaBarony{}

func init() {
    f := BTopana托帕纳.(*托帕纳TopanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "topana",
		TitleName: "托帕纳",
		TitleCode: "b_topana",
	}
}
