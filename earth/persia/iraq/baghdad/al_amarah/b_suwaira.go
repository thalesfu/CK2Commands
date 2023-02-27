package al_amarah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏瓦拉SuwairaBarony struct {
	feud.BaseBarony
}

var BSuwaira苏瓦拉 feud.Barony = &苏瓦拉SuwairaBarony{}

func init() {
    f := BSuwaira苏瓦拉.(*苏瓦拉SuwairaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suwaira",
		TitleName: "苏瓦拉",
		TitleCode: "b_suwaira",
	}
}
