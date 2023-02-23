package mesembria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 墨森布里亚MesembriaBarony struct {
	feud.BaseBarony
}

var BMesembria墨森布里亚 feud.Barony = &墨森布里亚MesembriaBarony{}

func init() {
	f := BMesembria墨森布里亚.(*墨森布里亚MesembriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mesembria",
		TitleName: "墨森布里亚",
		TitleCode: "b_mesembria",
	}
}
