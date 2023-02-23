package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉尔德库赫GerdkuhBarony struct {
	feud.BaseBarony
}

var BGerdkuh吉尔德库赫 feud.Barony = &吉尔德库赫GerdkuhBarony{}

func init() {
	f := BGerdkuh吉尔德库赫.(*吉尔德库赫GerdkuhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerdkuh",
		TitleName: "吉尔德库赫",
		TitleCode: "b_gerdkuh",
	}
}
