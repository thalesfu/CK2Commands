package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆布阿基OumelbouaghiBarony struct {
	feud.BaseBarony
}

var BOumelbouaghi乌姆布阿基 feud.Barony = &乌姆布阿基OumelbouaghiBarony{}

func init() {
    f := BOumelbouaghi乌姆布阿基.(*乌姆布阿基OumelbouaghiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oumelbouaghi",
		TitleName: "乌姆布阿基",
		TitleCode: "b_oumelbouaghi",
	}
}
