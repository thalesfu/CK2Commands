package guzgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿希纳弗堡QasrahnafBarony struct {
	feud.BaseBarony
}

var BQasrahnaf阿希纳弗堡 feud.Barony = &阿希纳弗堡QasrahnafBarony{}

func init() {
	f := BQasrahnaf阿希纳弗堡.(*阿希纳弗堡QasrahnafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qasrahnaf",
		TitleName: "阿希纳弗堡",
		TitleCode: "b_qasrahnaf",
	}
}
