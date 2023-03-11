package kyzyl-su

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克达什AkdashBarony struct {
	feud.BaseBarony
}

var BAkdash阿克达什 feud.Barony = &阿克达什AkdashBarony{}

func init() {
    f := BAkdash阿克达什.(*阿克达什AkdashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "akdash",
		TitleName: "阿克达什",
		TitleCode: "b_akdash",
	}
}
