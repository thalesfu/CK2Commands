package suf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尼迪内特Beni_dinhetBarony struct {
	feud.BaseBarony
}

var BBeni_dinhet贝尼迪内特 feud.Barony = &贝尼迪内特Beni_dinhetBarony{}

func init() {
    f := BBeni_dinhet贝尼迪内特.(*贝尼迪内特Beni_dinhetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beni_dinhet",
		TitleName: "贝尼迪内特",
		TitleCode: "b_beni_dinhet",
	}
}
