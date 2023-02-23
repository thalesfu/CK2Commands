package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达哈尔DadharBarony struct {
	feud.BaseBarony
}

var BDadhar达哈尔 feud.Barony = &达哈尔DadharBarony{}

func init() {
	f := BDadhar达哈尔.(*达哈尔DadharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dadhar",
		TitleName: "达哈尔",
		TitleCode: "b_dadhar",
	}
}
