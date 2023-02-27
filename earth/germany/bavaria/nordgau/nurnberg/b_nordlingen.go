package nurnberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 讷德林根NordlingenBarony struct {
	feud.BaseBarony
}

var BNordlingen讷德林根 feud.Barony = &讷德林根NordlingenBarony{}

func init() {
    f := BNordlingen讷德林根.(*讷德林根NordlingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nordlingen",
		TitleName: "讷德林根",
		TitleCode: "b_nordlingen",
	}
}
