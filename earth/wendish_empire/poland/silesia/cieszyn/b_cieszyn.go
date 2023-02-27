package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切申CieszynBarony struct {
	feud.BaseBarony
}

var BCieszyn切申 feud.Barony = &切申CieszynBarony{}

func init() {
    f := BCieszyn切申.(*切申CieszynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cieszyn",
		TitleName: "切申",
		TitleCode: "b_cieszyn",
	}
}
