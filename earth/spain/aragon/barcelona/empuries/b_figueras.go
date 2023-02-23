package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲格雷斯FiguerasBarony struct {
	feud.BaseBarony
}

var BFigueras菲格雷斯 feud.Barony = &菲格雷斯FiguerasBarony{}

func init() {
	f := BFigueras菲格雷斯.(*菲格雷斯FiguerasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "figueras",
		TitleName: "菲格雷斯",
		TitleCode: "b_figueras",
	}
}
