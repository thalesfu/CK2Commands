package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝德拉宾BedrabineBarony struct {
	feud.BaseBarony
}

var BBedrabine贝德拉宾 feud.Barony = &贝德拉宾BedrabineBarony{}

func init() {
    f := BBedrabine贝德拉宾.(*贝德拉宾BedrabineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bedrabine",
		TitleName: "贝德拉宾",
		TitleCode: "b_bedrabine",
	}
}
