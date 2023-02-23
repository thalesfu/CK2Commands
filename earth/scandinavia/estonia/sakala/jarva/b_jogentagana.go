package jarva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约根塔加纳JogentaganaBarony struct {
	feud.BaseBarony
}

var BJogentagana约根塔加纳 feud.Barony = &约根塔加纳JogentaganaBarony{}

func init() {
	f := BJogentagana约根塔加纳.(*约根塔加纳JogentaganaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jogentagana",
		TitleName: "约根塔加纳",
		TitleCode: "b_jogentagana",
	}
}
