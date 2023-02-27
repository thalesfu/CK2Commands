package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 堕罗三慕达罗DwarasamudraBarony struct {
	feud.BaseBarony
}

var BDwarasamudra堕罗三慕达罗 feud.Barony = &堕罗三慕达罗DwarasamudraBarony{}

func init() {
    f := BDwarasamudra堕罗三慕达罗.(*堕罗三慕达罗DwarasamudraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dwarasamudra",
		TitleName: "堕罗三慕达罗",
		TitleCode: "b_dwarasamudra",
	}
}
