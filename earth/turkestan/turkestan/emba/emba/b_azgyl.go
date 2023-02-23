package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿兹格尔AzgylBarony struct {
	feud.BaseBarony
}

var BAzgyl阿兹格尔 feud.Barony = &阿兹格尔AzgylBarony{}

func init() {
	f := BAzgyl阿兹格尔.(*阿兹格尔AzgylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azgyl",
		TitleName: "阿兹格尔",
		TitleCode: "b_azgyl",
	}
}
