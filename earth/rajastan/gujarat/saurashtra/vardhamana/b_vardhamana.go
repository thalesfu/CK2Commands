package vardhamana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 筏陀摩那VardhamanaBarony struct {
	feud.BaseBarony
}

var BVardhamana筏陀摩那 feud.Barony = &筏陀摩那VardhamanaBarony{}

func init() {
	f := BVardhamana筏陀摩那.(*筏陀摩那VardhamanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vardhamana",
		TitleName: "筏陀摩那",
		TitleCode: "b_vardhamana",
	}
}
