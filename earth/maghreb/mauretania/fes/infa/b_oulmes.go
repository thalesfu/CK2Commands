package infa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌勒马斯OulmesBarony struct {
	feud.BaseBarony
}

var BOulmes乌勒马斯 feud.Barony = &乌勒马斯OulmesBarony{}

func init() {
    f := BOulmes乌勒马斯.(*乌勒马斯OulmesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oulmes",
		TitleName: "乌勒马斯",
		TitleCode: "b_oulmes",
	}
}
