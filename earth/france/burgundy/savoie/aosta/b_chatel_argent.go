package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙泰尔阿尔让Chatel_argentBarony struct {
	feud.BaseBarony
}

var BChatel_argent沙泰尔阿尔让 feud.Barony = &沙泰尔阿尔让Chatel_argentBarony{}

func init() {
    f := BChatel_argent沙泰尔阿尔让.(*沙泰尔阿尔让Chatel_argentBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatel_argent",
		TitleName: "沙泰尔阿尔让",
		TitleCode: "b_chatel_argent",
	}
}
