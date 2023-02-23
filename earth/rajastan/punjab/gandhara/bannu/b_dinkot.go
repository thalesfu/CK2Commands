package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓科特DinkotBarony struct {
	feud.BaseBarony
}

var BDinkot邓科特 feud.Barony = &邓科特DinkotBarony{}

func init() {
	f := BDinkot邓科特.(*邓科特DinkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dinkot",
		TitleName: "邓科特",
		TitleCode: "b_dinkot",
	}
}
