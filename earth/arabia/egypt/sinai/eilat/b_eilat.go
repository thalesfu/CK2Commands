package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃拉特EilatBarony struct {
	feud.BaseBarony
}

var BEilat埃拉特 feud.Barony = &埃拉特EilatBarony{}

func init() {
	f := BEilat埃拉特.(*埃拉特EilatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eilat",
		TitleName: "埃拉特",
		TitleCode: "b_eilat",
	}
}
