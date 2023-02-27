package puri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾戈达JaugadaBarony struct {
	feud.BaseBarony
}

var BJaugada贾戈达 feud.Barony = &贾戈达JaugadaBarony{}

func init() {
    f := BJaugada贾戈达.(*贾戈达JaugadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jaugada",
		TitleName: "贾戈达",
		TitleCode: "b_jaugada",
	}
}
