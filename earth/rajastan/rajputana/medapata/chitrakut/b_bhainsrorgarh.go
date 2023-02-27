package chitrakut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜恩斯罗加尔BhainsrorgarhBarony struct {
	feud.BaseBarony
}

var BBhainsrorgarh拜恩斯罗加尔 feud.Barony = &拜恩斯罗加尔BhainsrorgarhBarony{}

func init() {
    f := BBhainsrorgarh拜恩斯罗加尔.(*拜恩斯罗加尔BhainsrorgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhainsrorgarh",
		TitleName: "拜恩斯罗加尔",
		TitleCode: "b_bhainsrorgarh",
	}
}
