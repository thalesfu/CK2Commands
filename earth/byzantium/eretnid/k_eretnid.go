package eretnid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EretnidKingdom interface {
	feud.Kingdom
}

type 埃雷特纳EretnidKingdom struct {
	feud.BaseKingdom
}

var KEretnid埃雷特纳 EretnidKingdom = &埃雷特纳EretnidKingdom{}

func init() {
	f := KEretnid埃雷特纳.(*埃雷特纳EretnidKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "eretnid",
		TitleName: "埃雷特纳",
		TitleCode: "k_eretnid",
		Dukes:     map[string]feud.Duke{},
	}

}
