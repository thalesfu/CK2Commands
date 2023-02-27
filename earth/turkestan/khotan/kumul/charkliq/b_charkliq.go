package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡克里克CharkliqBarony struct {
	feud.BaseBarony
}

var BCharkliq卡克里克 feud.Barony = &卡克里克CharkliqBarony{}

func init() {
    f := BCharkliq卡克里克.(*卡克里克CharkliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "charkliq",
		TitleName: "卡克里克",
		TitleCode: "b_charkliq",
	}
}
