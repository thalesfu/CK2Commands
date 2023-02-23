package genoa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福斯迪诺沃FosdinovoBarony struct {
	feud.BaseBarony
}

var BFosdinovo福斯迪诺沃 feud.Barony = &福斯迪诺沃FosdinovoBarony{}

func init() {
	f := BFosdinovo福斯迪诺沃.(*福斯迪诺沃FosdinovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fosdinovo",
		TitleName: "福斯迪诺沃",
		TitleCode: "b_fosdinovo",
	}
}
