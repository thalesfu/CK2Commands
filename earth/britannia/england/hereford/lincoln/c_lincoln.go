package lincoln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LincolnCounty interface {
	feud.County
	BBardney巴德尼() feud.Barony
	BBoston波士顿() feud.Barony
	BGainsborough盖恩斯伯勒() feud.Barony
	BGrantham格兰瑟姆() feud.Barony
	BLincoln林肯() feud.Barony
	BLouth劳斯() feud.Barony
	BSpalding斯伯丁() feud.Barony
	BStamford斯坦福() feud.Barony
}

type 林肯LincolnCounty struct {
	feud.BaseCounty
	巴德尼Bardney        feud.Barony
	波士顿Boston         feud.Barony
	盖恩斯伯勒Gainsborough feud.Barony
	格兰瑟姆Grantham      feud.Barony
	林肯Lincoln         feud.Barony
	劳斯Louth           feud.Barony
	斯伯丁Spalding       feud.Barony
	斯坦福Stamford       feud.Barony
}

func (c *林肯LincolnCounty) BBardney巴德尼() feud.Barony {
	return c.巴德尼Bardney
}

func (c *林肯LincolnCounty) BBoston波士顿() feud.Barony {
	return c.波士顿Boston
}

func (c *林肯LincolnCounty) BGainsborough盖恩斯伯勒() feud.Barony {
	return c.盖恩斯伯勒Gainsborough
}

func (c *林肯LincolnCounty) BGrantham格兰瑟姆() feud.Barony {
	return c.格兰瑟姆Grantham
}

func (c *林肯LincolnCounty) BLincoln林肯() feud.Barony {
	return c.林肯Lincoln
}

func (c *林肯LincolnCounty) BLouth劳斯() feud.Barony {
	return c.劳斯Louth
}

func (c *林肯LincolnCounty) BSpalding斯伯丁() feud.Barony {
	return c.斯伯丁Spalding
}

func (c *林肯LincolnCounty) BStamford斯坦福() feud.Barony {
	return c.斯坦福Stamford
}

var CLincoln林肯 LincolnCounty = &林肯LincolnCounty{}

func init() {
	f := CLincoln林肯.(*林肯LincolnCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "61",
		Title:     "lincoln",
		TitleName: "林肯",
		TitleCode: "c_lincoln",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴德尼Bardney = BBardney巴德尼
	f.巴德尼Bardney.SetParent(f)

	f.波士顿Boston = BBoston波士顿
	f.波士顿Boston.SetParent(f)

	f.盖恩斯伯勒Gainsborough = BGainsborough盖恩斯伯勒
	f.盖恩斯伯勒Gainsborough.SetParent(f)

	f.格兰瑟姆Grantham = BGrantham格兰瑟姆
	f.格兰瑟姆Grantham.SetParent(f)

	f.林肯Lincoln = BLincoln林肯
	f.林肯Lincoln.SetParent(f)

	f.劳斯Louth = BLouth劳斯
	f.劳斯Louth.SetParent(f)

	f.斯伯丁Spalding = BSpalding斯伯丁
	f.斯伯丁Spalding.SetParent(f)

	f.斯坦福Stamford = BStamford斯坦福
	f.斯坦福Stamford.SetParent(f)

}
