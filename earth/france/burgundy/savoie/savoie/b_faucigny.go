package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福西尼FaucignyBarony struct {
	feud.BaseBarony
}

var BFaucigny福西尼 feud.Barony = &福西尼FaucignyBarony{}

func init() {
	f := BFaucigny福西尼.(*福西尼FaucignyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "faucigny",
		TitleName: "福西尼",
		TitleCode: "b_faucigny",
	}
}
