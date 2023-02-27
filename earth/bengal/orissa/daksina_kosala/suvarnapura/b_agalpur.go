package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阏伽罗补罗AgalpurBarony struct {
	feud.BaseBarony
}

var BAgalpur阏伽罗补罗 feud.Barony = &阏伽罗补罗AgalpurBarony{}

func init() {
    f := BAgalpur阏伽罗补罗.(*阏伽罗补罗AgalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agalpur",
		TitleName: "阏伽罗补罗",
		TitleCode: "b_agalpur",
	}
}
