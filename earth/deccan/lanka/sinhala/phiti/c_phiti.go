package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PhitiCounty interface {
	feud.County
	BAnuradhapura阿㝹罗陀补罗() feud.Barony
	BAtamasthana阿吒摩萨傥那() feud.Barony
	BMahatittha摩诃迭他() feud.Barony
	BMannara曼那罗() feud.Barony
	BMihintale密亨达勒() feud.Barony
	BPunkhagama封佉伽摩() feud.Barony
	BSubhagiri须婆耆厘() feud.Barony
}

type 比醯帝PhitiCounty struct {
	feud.BaseCounty
	阿㝹罗陀补罗Anuradhapura feud.Barony
	阿吒摩萨傥那Atamasthana  feud.Barony
	摩诃迭他Mahatittha     feud.Barony
	曼那罗Mannara         feud.Barony
	密亨达勒Mihintale      feud.Barony
	封佉伽摩Punkhagama     feud.Barony
	须婆耆厘Subhagiri      feud.Barony
}

func (c *比醯帝PhitiCounty) BAnuradhapura阿㝹罗陀补罗() feud.Barony {
	return c.阿㝹罗陀补罗Anuradhapura
}

func (c *比醯帝PhitiCounty) BAtamasthana阿吒摩萨傥那() feud.Barony {
	return c.阿吒摩萨傥那Atamasthana
}

func (c *比醯帝PhitiCounty) BMahatittha摩诃迭他() feud.Barony {
	return c.摩诃迭他Mahatittha
}

func (c *比醯帝PhitiCounty) BMannara曼那罗() feud.Barony {
	return c.曼那罗Mannara
}

func (c *比醯帝PhitiCounty) BMihintale密亨达勒() feud.Barony {
	return c.密亨达勒Mihintale
}

func (c *比醯帝PhitiCounty) BPunkhagama封佉伽摩() feud.Barony {
	return c.封佉伽摩Punkhagama
}

func (c *比醯帝PhitiCounty) BSubhagiri须婆耆厘() feud.Barony {
	return c.须婆耆厘Subhagiri
}

var CPhiti比醯帝 PhitiCounty = &比醯帝PhitiCounty{}

func init() {
	f := CPhiti比醯帝.(*比醯帝PhitiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1195",
		Title:     "phiti",
		TitleName: "比醯帝",
		TitleCode: "c_phiti",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿㝹罗陀补罗Anuradhapura = BAnuradhapura阿㝹罗陀补罗
	f.阿㝹罗陀补罗Anuradhapura.SetParent(f)

	f.阿吒摩萨傥那Atamasthana = BAtamasthana阿吒摩萨傥那
	f.阿吒摩萨傥那Atamasthana.SetParent(f)

	f.摩诃迭他Mahatittha = BMahatittha摩诃迭他
	f.摩诃迭他Mahatittha.SetParent(f)

	f.曼那罗Mannara = BMannara曼那罗
	f.曼那罗Mannara.SetParent(f)

	f.密亨达勒Mihintale = BMihintale密亨达勒
	f.密亨达勒Mihintale.SetParent(f)

	f.封佉伽摩Punkhagama = BPunkhagama封佉伽摩
	f.封佉伽摩Punkhagama.SetParent(f)

	f.须婆耆厘Subhagiri = BSubhagiri须婆耆厘
	f.须婆耆厘Subhagiri.SetParent(f)

}
