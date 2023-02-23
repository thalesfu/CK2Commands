package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米歇尔多夫MicheldorfBarony struct {
	feud.BaseBarony
}

var BMicheldorf米歇尔多夫 feud.Barony = &米歇尔多夫MicheldorfBarony{}

func init() {
	f := BMicheldorf米歇尔多夫.(*米歇尔多夫MicheldorfBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "micheldorf",
		TitleName: "米歇尔多夫",
		TitleCode: "b_micheldorf",
	}
}
