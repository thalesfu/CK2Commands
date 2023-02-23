package sjaelland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯隆堡KalundborgBarony struct {
	feud.BaseBarony
}

var BKalundborg凯隆堡 feud.Barony = &凯隆堡KalundborgBarony{}

func init() {
	f := BKalundborg凯隆堡.(*凯隆堡KalundborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalundborg",
		TitleName: "凯隆堡",
		TitleCode: "b_kalundborg",
	}
}
