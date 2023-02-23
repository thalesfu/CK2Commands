package kutch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 侯伽补罗HukampuraBarony struct {
	feud.BaseBarony
}

var BHukampura侯伽补罗 feud.Barony = &侯伽补罗HukampuraBarony{}

func init() {
	f := BHukampura侯伽补罗.(*侯伽补罗HukampuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hukampura",
		TitleName: "侯伽补罗",
		TitleCode: "b_hukampura",
	}
}
