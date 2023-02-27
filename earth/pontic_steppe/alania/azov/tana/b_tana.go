package tana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔纳TanaBarony struct {
	feud.BaseBarony
}

var BTana塔纳 feud.Barony = &塔纳TanaBarony{}

func init() {
    f := BTana塔纳.(*塔纳TanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tana",
		TitleName: "塔纳",
		TitleCode: "b_tana",
	}
}
