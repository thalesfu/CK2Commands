package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MagadhaCounty interface {
    feud.County
    BBarh婆罗() 	feud.Barony
    BBihar_sharif毗诃罗谢里夫() 	feud.Barony
    BManer摩内罗() 	feud.Barony
    BOdantapuri飞行寺() 	feud.Barony
    BPataliputra华氏城() 	feud.Barony
    BRajagrha王舍城() 	feud.Barony
    BTriveni帝力吠尼() 	feud.Barony
}

type 摩揭陀MagadhaCounty struct {
	feud.BaseCounty
	婆罗Barh 	feud.Barony
	毗诃罗谢里夫Bihar_sharif 	feud.Barony
	摩内罗Maner 	feud.Barony
	飞行寺Odantapuri 	feud.Barony
	华氏城Pataliputra 	feud.Barony
	王舍城Rajagrha 	feud.Barony
	帝力吠尼Triveni 	feud.Barony
}

func (c *摩揭陀MagadhaCounty) BBarh婆罗() feud.Barony {
	return c.婆罗Barh
}
    
func (c *摩揭陀MagadhaCounty) BBihar_sharif毗诃罗谢里夫() feud.Barony {
	return c.毗诃罗谢里夫Bihar_sharif
}
    
func (c *摩揭陀MagadhaCounty) BManer摩内罗() feud.Barony {
	return c.摩内罗Maner
}
    
func (c *摩揭陀MagadhaCounty) BOdantapuri飞行寺() feud.Barony {
	return c.飞行寺Odantapuri
}
    
func (c *摩揭陀MagadhaCounty) BPataliputra华氏城() feud.Barony {
	return c.华氏城Pataliputra
}
    
func (c *摩揭陀MagadhaCounty) BRajagrha王舍城() feud.Barony {
	return c.王舍城Rajagrha
}
    
func (c *摩揭陀MagadhaCounty) BTriveni帝力吠尼() feud.Barony {
	return c.帝力吠尼Triveni
}
    
var CMagadha摩揭陀 MagadhaCounty = &摩揭陀MagadhaCounty{}

func init() {
	f := CMagadha摩揭陀.(*摩揭陀MagadhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1154",
		Title:     "magadha",
		TitleName: "摩揭陀",
		TitleCode: "c_magadha",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆罗Barh = BBarh婆罗
	f.婆罗Barh.SetParent(f)
	
	f.毗诃罗谢里夫Bihar_sharif = BBihar_sharif毗诃罗谢里夫
	f.毗诃罗谢里夫Bihar_sharif.SetParent(f)
	
	f.摩内罗Maner = BManer摩内罗
	f.摩内罗Maner.SetParent(f)
	
	f.飞行寺Odantapuri = BOdantapuri飞行寺
	f.飞行寺Odantapuri.SetParent(f)
	
	f.华氏城Pataliputra = BPataliputra华氏城
	f.华氏城Pataliputra.SetParent(f)
	
	f.王舍城Rajagrha = BRajagrha王舍城
	f.王舍城Rajagrha.SetParent(f)
	
	f.帝力吠尼Triveni = BTriveni帝力吠尼
	f.帝力吠尼Triveni.SetParent(f)
	
}
