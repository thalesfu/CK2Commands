package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哥达GothaBarony struct {
	feud.BaseBarony
}

var BGotha哥达 feud.Barony = &哥达GothaBarony{}

func init() {
	f := BGotha哥达.(*哥达GothaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gotha",
		TitleName: "哥达",
		TitleCode: "b_gotha",
	}
}
