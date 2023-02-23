package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SomersetCounty interface {
	feud.County
	BBridgwater布里奇沃特() feud.Barony
	BGlastonbury格拉斯顿伯里() feud.Barony
	BIlchester伊尔切斯特() feud.Barony
	BMuchelney马乔尼() feud.Barony
	BTaunton汤顿() feud.Barony
	BYeovil约维尔() feud.Barony
}

type 萨默塞特SomersetCounty struct {
	feud.BaseCounty
	布里奇沃特Bridgwater   feud.Barony
	格拉斯顿伯里Glastonbury feud.Barony
	伊尔切斯特Ilchester    feud.Barony
	马乔尼Muchelney      feud.Barony
	汤顿Taunton         feud.Barony
	约维尔Yeovil         feud.Barony
}

func (c *萨默塞特SomersetCounty) BBridgwater布里奇沃特() feud.Barony {
	return c.布里奇沃特Bridgwater
}

func (c *萨默塞特SomersetCounty) BGlastonbury格拉斯顿伯里() feud.Barony {
	return c.格拉斯顿伯里Glastonbury
}

func (c *萨默塞特SomersetCounty) BIlchester伊尔切斯特() feud.Barony {
	return c.伊尔切斯特Ilchester
}

func (c *萨默塞特SomersetCounty) BMuchelney马乔尼() feud.Barony {
	return c.马乔尼Muchelney
}

func (c *萨默塞特SomersetCounty) BTaunton汤顿() feud.Barony {
	return c.汤顿Taunton
}

func (c *萨默塞特SomersetCounty) BYeovil约维尔() feud.Barony {
	return c.约维尔Yeovil
}

var CSomerset萨默塞特 SomersetCounty = &萨默塞特SomersetCounty{}

func init() {
	f := CSomerset萨默塞特.(*萨默塞特SomersetCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "28",
		Title:     "somerset",
		TitleName: "萨默塞特",
		TitleCode: "c_somerset",
		Baronies:  map[string]feud.Barony{},
	}

	f.布里奇沃特Bridgwater = BBridgwater布里奇沃特
	f.布里奇沃特Bridgwater.SetParent(f)

	f.格拉斯顿伯里Glastonbury = BGlastonbury格拉斯顿伯里
	f.格拉斯顿伯里Glastonbury.SetParent(f)

	f.伊尔切斯特Ilchester = BIlchester伊尔切斯特
	f.伊尔切斯特Ilchester.SetParent(f)

	f.马乔尼Muchelney = BMuchelney马乔尼
	f.马乔尼Muchelney.SetParent(f)

	f.汤顿Taunton = BTaunton汤顿
	f.汤顿Taunton.SetParent(f)

	f.约维尔Yeovil = BYeovil约维尔
	f.约维尔Yeovil.SetParent(f)

}
