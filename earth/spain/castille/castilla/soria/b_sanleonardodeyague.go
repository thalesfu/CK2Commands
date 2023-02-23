package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣莱昂纳多德亚圭SanleonardodeyagueBarony struct {
	feud.BaseBarony
}

var BSanleonardodeyague圣莱昂纳多德亚圭 feud.Barony = &圣莱昂纳多德亚圭SanleonardodeyagueBarony{}

func init() {
	f := BSanleonardodeyague圣莱昂纳多德亚圭.(*圣莱昂纳多德亚圭SanleonardodeyagueBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanleonardodeyague",
		TitleName: "圣莱昂纳多德亚圭",
		TitleCode: "b_sanleonardodeyague",
	}
}
