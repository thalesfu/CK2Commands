package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙泰洛德朗ChatelaudrenBarony struct {
	feud.BaseBarony
}

var BChatelaudren沙泰洛德朗 feud.Barony = &沙泰洛德朗ChatelaudrenBarony{}

func init() {
	f := BChatelaudren沙泰洛德朗.(*沙泰洛德朗ChatelaudrenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chatelaudren",
		TitleName: "沙泰洛德朗",
		TitleCode: "b_chatelaudren",
	}
}
