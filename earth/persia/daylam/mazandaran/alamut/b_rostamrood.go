package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗斯达摩罗RostamroodBarony struct {
	feud.BaseBarony
}

var BRostamrood罗斯达摩罗 feud.Barony = &罗斯达摩罗RostamroodBarony{}

func init() {
    f := BRostamrood罗斯达摩罗.(*罗斯达摩罗RostamroodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rostamrood",
		TitleName: "罗斯达摩罗",
		TitleCode: "b_rostamrood",
	}
}
