package angouleme

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉特朗沙德LatranchadeBarony struct {
	feud.BaseBarony
}

var BLatranchade拉特朗沙德 feud.Barony = &拉特朗沙德LatranchadeBarony{}

func init() {
	f := BLatranchade拉特朗沙德.(*拉特朗沙德LatranchadeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "latranchade",
		TitleName: "拉特朗沙德",
		TitleCode: "b_latranchade",
	}
}
