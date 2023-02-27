package tell_bashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼比季ManbijBarony struct {
	feud.BaseBarony
}

var BManbij曼比季 feud.Barony = &曼比季ManbijBarony{}

func init() {
    f := BManbij曼比季.(*曼比季ManbijBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manbij",
		TitleName: "曼比季",
		TitleCode: "b_manbij",
	}
}
