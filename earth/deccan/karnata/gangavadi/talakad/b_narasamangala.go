package talakad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那罗娑茫伽罗NarasamangalaBarony struct {
	feud.BaseBarony
}

var BNarasamangala那罗娑茫伽罗 feud.Barony = &那罗娑茫伽罗NarasamangalaBarony{}

func init() {
	f := BNarasamangala那罗娑茫伽罗.(*那罗娑茫伽罗NarasamangalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narasamangala",
		TitleName: "那罗娑茫伽罗",
		TitleCode: "b_narasamangala",
	}
}
