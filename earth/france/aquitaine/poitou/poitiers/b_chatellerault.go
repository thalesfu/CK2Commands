package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙泰勒罗ChatelleraultBarony struct {
	feud.BaseBarony
}

var BChatellerault沙泰勒罗 feud.Barony = &沙泰勒罗ChatelleraultBarony{}

func init() {
    f := BChatellerault沙泰勒罗.(*沙泰勒罗ChatelleraultBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatellerault",
		TitleName: "沙泰勒罗",
		TitleCode: "b_chatellerault",
	}
}
