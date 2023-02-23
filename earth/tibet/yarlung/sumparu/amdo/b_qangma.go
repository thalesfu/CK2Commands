package amdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 强玛QangmaBarony struct {
	feud.BaseBarony
}

var BQangma强玛 feud.Barony = &强玛QangmaBarony{}

func init() {
	f := BQangma强玛.(*强玛QangmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qangma",
		TitleName: "强玛",
		TitleCode: "b_qangma",
	}
}
