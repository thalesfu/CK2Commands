package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马贾阿卡比尔MajaralkabirBarony struct {
	feud.BaseBarony
}

var BMajaralkabir马贾阿卡比尔 feud.Barony = &马贾阿卡比尔MajaralkabirBarony{}

func init() {
    f := BMajaralkabir马贾阿卡比尔.(*马贾阿卡比尔MajaralkabirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "majaralkabir",
		TitleName: "马贾阿卡比尔",
		TitleCode: "b_majaralkabir",
	}
}
