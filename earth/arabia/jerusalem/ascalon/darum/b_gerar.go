package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基拉耳GerarBarony struct {
	feud.BaseBarony
}

var BGerar基拉耳 feud.Barony = &基拉耳GerarBarony{}

func init() {
	f := BGerar基拉耳.(*基拉耳GerarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerar",
		TitleName: "基拉耳",
		TitleCode: "b_gerar",
	}
}
