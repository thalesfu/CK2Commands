package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ReniCounty interface {
	feud.County
	BBidasar毗陀娑() feud.Barony
	BDadrewa陀多瓦() feud.Barony
	BMandir寺庙() feud.Barony
	BNiralapalli尼罗罗波梨() feud.Barony
	BPalshet波湿怛() feud.Barony
	BReni梨尼() feud.Barony
	BSidhmukh悉陀目佉() feud.Barony
}

type 梨尼ReniCounty struct {
	feud.BaseCounty
	毗陀娑Bidasar       feud.Barony
	陀多瓦Dadrewa       feud.Barony
	寺庙Mandir         feud.Barony
	尼罗罗波梨Niralapalli feud.Barony
	波湿怛Palshet       feud.Barony
	梨尼Reni           feud.Barony
	悉陀目佉Sidhmukh     feud.Barony
}

func (c *梨尼ReniCounty) BBidasar毗陀娑() feud.Barony {
	return c.毗陀娑Bidasar
}

func (c *梨尼ReniCounty) BDadrewa陀多瓦() feud.Barony {
	return c.陀多瓦Dadrewa
}

func (c *梨尼ReniCounty) BMandir寺庙() feud.Barony {
	return c.寺庙Mandir
}

func (c *梨尼ReniCounty) BNiralapalli尼罗罗波梨() feud.Barony {
	return c.尼罗罗波梨Niralapalli
}

func (c *梨尼ReniCounty) BPalshet波湿怛() feud.Barony {
	return c.波湿怛Palshet
}

func (c *梨尼ReniCounty) BReni梨尼() feud.Barony {
	return c.梨尼Reni
}

func (c *梨尼ReniCounty) BSidhmukh悉陀目佉() feud.Barony {
	return c.悉陀目佉Sidhmukh
}

var CReni梨尼 ReniCounty = &梨尼ReniCounty{}

func init() {
	f := CReni梨尼.(*梨尼ReniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1350",
		Title:     "reni",
		TitleName: "梨尼",
		TitleCode: "c_reni",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗陀娑Bidasar = BBidasar毗陀娑
	f.毗陀娑Bidasar.SetParent(f)

	f.陀多瓦Dadrewa = BDadrewa陀多瓦
	f.陀多瓦Dadrewa.SetParent(f)

	f.寺庙Mandir = BMandir寺庙
	f.寺庙Mandir.SetParent(f)

	f.尼罗罗波梨Niralapalli = BNiralapalli尼罗罗波梨
	f.尼罗罗波梨Niralapalli.SetParent(f)

	f.波湿怛Palshet = BPalshet波湿怛
	f.波湿怛Palshet.SetParent(f)

	f.梨尼Reni = BReni梨尼
	f.梨尼Reni.SetParent(f)

	f.悉陀目佉Sidhmukh = BSidhmukh悉陀目佉
	f.悉陀目佉Sidhmukh.SetParent(f)

}
