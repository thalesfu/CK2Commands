package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乃东NedongBarony struct {
	feud.BaseBarony
}

var BNedong乃东 feud.Barony = &乃东NedongBarony{}

func init() {
	f := BNedong乃东.(*乃东NedongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nedong",
		TitleName: "乃东",
		TitleCode: "b_nedong",
	}
}
