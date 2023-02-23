package gao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加奥GaoBarony struct {
	feud.BaseBarony
}

var BGao加奥 feud.Barony = &加奥GaoBarony{}

func init() {
	f := BGao加奥.(*加奥GaoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gao",
		TitleName: "加奥",
		TitleCode: "b_gao",
	}
}
