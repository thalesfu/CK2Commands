package candradvipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CandradvipaCounty interface {
	feud.County
	BBagerhat巴格哈特() feud.Barony
	BBakarganj波迦罗犍阇() feud.Barony
	BBilkhet毗吉多() feud.Barony
	BCandradvipa月洲() feud.Barony
	BDattapulia陀怛多补梨() feud.Barony
	BKhulna丘罗那() feud.Barony
	BSathkira萨诃吉罗() feud.Barony
}

type 月洲CandradvipaCounty struct {
	feud.BaseCounty
	巴格哈特Bagerhat    feud.Barony
	波迦罗犍阇Bakarganj  feud.Barony
	毗吉多Bilkhet      feud.Barony
	月洲Candradvipa   feud.Barony
	陀怛多补梨Dattapulia feud.Barony
	丘罗那Khulna       feud.Barony
	萨诃吉罗Sathkira    feud.Barony
}

func (c *月洲CandradvipaCounty) BBagerhat巴格哈特() feud.Barony {
	return c.巴格哈特Bagerhat
}

func (c *月洲CandradvipaCounty) BBakarganj波迦罗犍阇() feud.Barony {
	return c.波迦罗犍阇Bakarganj
}

func (c *月洲CandradvipaCounty) BBilkhet毗吉多() feud.Barony {
	return c.毗吉多Bilkhet
}

func (c *月洲CandradvipaCounty) BCandradvipa月洲() feud.Barony {
	return c.月洲Candradvipa
}

func (c *月洲CandradvipaCounty) BDattapulia陀怛多补梨() feud.Barony {
	return c.陀怛多补梨Dattapulia
}

func (c *月洲CandradvipaCounty) BKhulna丘罗那() feud.Barony {
	return c.丘罗那Khulna
}

func (c *月洲CandradvipaCounty) BSathkira萨诃吉罗() feud.Barony {
	return c.萨诃吉罗Sathkira
}

var CCandradvipa月洲 CandradvipaCounty = &月洲CandradvipaCounty{}

func init() {
	f := CCandradvipa月洲.(*月洲CandradvipaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1236",
		Title:     "candradvipa",
		TitleName: "月洲",
		TitleCode: "c_candradvipa",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴格哈特Bagerhat = BBagerhat巴格哈特
	f.巴格哈特Bagerhat.SetParent(f)

	f.波迦罗犍阇Bakarganj = BBakarganj波迦罗犍阇
	f.波迦罗犍阇Bakarganj.SetParent(f)

	f.毗吉多Bilkhet = BBilkhet毗吉多
	f.毗吉多Bilkhet.SetParent(f)

	f.月洲Candradvipa = BCandradvipa月洲
	f.月洲Candradvipa.SetParent(f)

	f.陀怛多补梨Dattapulia = BDattapulia陀怛多补梨
	f.陀怛多补梨Dattapulia.SetParent(f)

	f.丘罗那Khulna = BKhulna丘罗那
	f.丘罗那Khulna.SetParent(f)

	f.萨诃吉罗Sathkira = BSathkira萨诃吉罗
	f.萨诃吉罗Sathkira.SetParent(f)

}
