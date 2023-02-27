package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 信度梨耶HindoliaBarony struct {
	feud.BaseBarony
}

var BHindolia信度梨耶 feud.Barony = &信度梨耶HindoliaBarony{}

func init() {
    f := BHindolia信度梨耶.(*信度梨耶HindoliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hindolia",
		TitleName: "信度梨耶",
		TitleCode: "b_hindolia",
	}
}
