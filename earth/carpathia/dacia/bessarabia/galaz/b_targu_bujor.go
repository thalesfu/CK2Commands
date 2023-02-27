package galaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔古布若尔Targu_bujorBarony struct {
	feud.BaseBarony
}

var BTargu_bujor特尔古布若尔 feud.Barony = &特尔古布若尔Targu_bujorBarony{}

func init() {
    f := BTargu_bujor特尔古布若尔.(*特尔古布若尔Targu_bujorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "targu_bujor",
		TitleName: "特尔古布若尔",
		TitleCode: "b_targu_bujor",
	}
}
