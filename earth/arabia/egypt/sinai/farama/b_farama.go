package farama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀卢西亚FaramaBarony struct {
	feud.BaseBarony
}

var BFarama珀卢西亚 feud.Barony = &珀卢西亚FaramaBarony{}

func init() {
	f := BFarama珀卢西亚.(*珀卢西亚FaramaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farama",
		TitleName: "珀卢西亚",
		TitleCode: "b_farama",
	}
}
