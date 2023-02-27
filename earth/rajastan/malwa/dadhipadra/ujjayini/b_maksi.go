package ujjayini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩克西MaksiBarony struct {
	feud.BaseBarony
}

var BMaksi摩克西 feud.Barony = &摩克西MaksiBarony{}

func init() {
    f := BMaksi摩克西.(*摩克西MaksiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maksi",
		TitleName: "摩克西",
		TitleCode: "b_maksi",
	}
}
