package chunar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毗梨拘梨BilikoreBarony struct {
	feud.BaseBarony
}

var BBilikore毗梨拘梨 feud.Barony = &毗梨拘梨BilikoreBarony{}

func init() {
    f := BBilikore毗梨拘梨.(*毗梨拘梨BilikoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilikore",
		TitleName: "毗梨拘梨",
		TitleCode: "b_bilikore",
	}
}
