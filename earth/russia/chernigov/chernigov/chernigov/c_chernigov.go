package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChernigovCounty interface {
	feud.County
	BChernigov切尔尼戈夫() feud.Barony
	BHorodina戈罗迪纳() feud.Barony
	BKorop科罗普() feud.Barony
	BMena梅纳() feud.Barony
	BRipky里普基() feud.Barony
	BSnov斯诺夫() feud.Barony
	BSosnytsia索斯尼齐亚() feud.Barony
}

type 切尔尼戈夫ChernigovCounty struct {
	feud.BaseCounty
	切尔尼戈夫Chernigov feud.Barony
	戈罗迪纳Horodina   feud.Barony
	科罗普Korop       feud.Barony
	梅纳Mena         feud.Barony
	里普基Ripky       feud.Barony
	斯诺夫Snov        feud.Barony
	索斯尼齐亚Sosnytsia feud.Barony
}

func (c *切尔尼戈夫ChernigovCounty) BChernigov切尔尼戈夫() feud.Barony {
	return c.切尔尼戈夫Chernigov
}

func (c *切尔尼戈夫ChernigovCounty) BHorodina戈罗迪纳() feud.Barony {
	return c.戈罗迪纳Horodina
}

func (c *切尔尼戈夫ChernigovCounty) BKorop科罗普() feud.Barony {
	return c.科罗普Korop
}

func (c *切尔尼戈夫ChernigovCounty) BMena梅纳() feud.Barony {
	return c.梅纳Mena
}

func (c *切尔尼戈夫ChernigovCounty) BRipky里普基() feud.Barony {
	return c.里普基Ripky
}

func (c *切尔尼戈夫ChernigovCounty) BSnov斯诺夫() feud.Barony {
	return c.斯诺夫Snov
}

func (c *切尔尼戈夫ChernigovCounty) BSosnytsia索斯尼齐亚() feud.Barony {
	return c.索斯尼齐亚Sosnytsia
}

var CChernigov切尔尼戈夫 ChernigovCounty = &切尔尼戈夫ChernigovCounty{}

func init() {
	f := CChernigov切尔尼戈夫.(*切尔尼戈夫ChernigovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "554",
		Title:     "chernigov",
		TitleName: "切尔尼戈夫",
		TitleCode: "c_chernigov",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔尼戈夫Chernigov = BChernigov切尔尼戈夫
	f.切尔尼戈夫Chernigov.SetParent(f)

	f.戈罗迪纳Horodina = BHorodina戈罗迪纳
	f.戈罗迪纳Horodina.SetParent(f)

	f.科罗普Korop = BKorop科罗普
	f.科罗普Korop.SetParent(f)

	f.梅纳Mena = BMena梅纳
	f.梅纳Mena.SetParent(f)

	f.里普基Ripky = BRipky里普基
	f.里普基Ripky.SetParent(f)

	f.斯诺夫Snov = BSnov斯诺夫
	f.斯诺夫Snov.SetParent(f)

	f.索斯尼齐亚Sosnytsia = BSosnytsia索斯尼齐亚
	f.索斯尼齐亚Sosnytsia.SetParent(f)

}
