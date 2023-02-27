package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡达纳迈克尔Godana_mikaelBarony struct {
	feud.BaseBarony
}

var BGodana_mikael贡达纳迈克尔 feud.Barony = &贡达纳迈克尔Godana_mikaelBarony{}

func init() {
    f := BGodana_mikael贡达纳迈克尔.(*贡达纳迈克尔Godana_mikaelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "godana_mikael",
		TitleName: "贡达纳迈克尔",
		TitleCode: "b_godana_mikael",
	}
}
