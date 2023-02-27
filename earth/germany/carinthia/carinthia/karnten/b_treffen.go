package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特雷芬TreffenBarony struct {
	feud.BaseBarony
}

var BTreffen特雷芬 feud.Barony = &特雷芬TreffenBarony{}

func init() {
    f := BTreffen特雷芬.(*特雷芬TreffenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "treffen",
		TitleName: "特雷芬",
		TitleCode: "b_treffen",
	}
}
