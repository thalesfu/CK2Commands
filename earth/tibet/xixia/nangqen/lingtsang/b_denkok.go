package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓柯DenkokBarony struct {
	feud.BaseBarony
}

var BDenkok邓柯 feud.Barony = &邓柯DenkokBarony{}

func init() {
    f := BDenkok邓柯.(*邓柯DenkokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "denkok",
		TitleName: "邓柯",
		TitleCode: "b_denkok",
	}
}
