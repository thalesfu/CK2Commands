package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 腾格尔TengerBarony struct {
	feud.BaseBarony
}

var BTenger腾格尔 feud.Barony = &腾格尔TengerBarony{}

func init() {
	f := BTenger腾格尔.(*腾格尔TengerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tenger",
		TitleName: "腾格尔",
		TitleCode: "b_tenger",
	}
}
