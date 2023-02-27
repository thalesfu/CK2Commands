package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AostaCounty interface {
    feud.County
    BAosta奥斯塔() 	feud.Barony
    BAostachatillon沙蒂永() 	feud.Barony
    BArnad阿尔纳德() 	feud.Barony
    BChatel_argent沙泰尔阿尔让() 	feud.Barony
    BFenis费尼斯() 	feud.Barony
    BSt_jacqueme圣雅凯姆() 	feud.Barony
    BSt_ours圣乌尔() 	feud.Barony
    BSt_pierre圣皮埃尔() 	feud.Barony
}

type 奥斯塔AostaCounty struct {
	feud.BaseCounty
	奥斯塔Aosta 	feud.Barony
	沙蒂永Aostachatillon 	feud.Barony
	阿尔纳德Arnad 	feud.Barony
	沙泰尔阿尔让Chatel_argent 	feud.Barony
	费尼斯Fenis 	feud.Barony
	圣雅凯姆St_jacqueme 	feud.Barony
	圣乌尔St_ours 	feud.Barony
	圣皮埃尔St_pierre 	feud.Barony
}

func (c *奥斯塔AostaCounty) BAosta奥斯塔() feud.Barony {
	return c.奥斯塔Aosta
}
    
func (c *奥斯塔AostaCounty) BAostachatillon沙蒂永() feud.Barony {
	return c.沙蒂永Aostachatillon
}
    
func (c *奥斯塔AostaCounty) BArnad阿尔纳德() feud.Barony {
	return c.阿尔纳德Arnad
}
    
func (c *奥斯塔AostaCounty) BChatel_argent沙泰尔阿尔让() feud.Barony {
	return c.沙泰尔阿尔让Chatel_argent
}
    
func (c *奥斯塔AostaCounty) BFenis费尼斯() feud.Barony {
	return c.费尼斯Fenis
}
    
func (c *奥斯塔AostaCounty) BSt_jacqueme圣雅凯姆() feud.Barony {
	return c.圣雅凯姆St_jacqueme
}
    
func (c *奥斯塔AostaCounty) BSt_ours圣乌尔() feud.Barony {
	return c.圣乌尔St_ours
}
    
func (c *奥斯塔AostaCounty) BSt_pierre圣皮埃尔() feud.Barony {
	return c.圣皮埃尔St_pierre
}
    
var CAosta奥斯塔 AostaCounty = &奥斯塔AostaCounty{}

func init() {
	f := CAosta奥斯塔.(*奥斯塔AostaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1610",
		Title:     "aosta",
		TitleName: "奥斯塔",
		TitleCode: "c_aosta",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥斯塔Aosta = BAosta奥斯塔
	f.奥斯塔Aosta.SetParent(f)
	
	f.沙蒂永Aostachatillon = BAostachatillon沙蒂永
	f.沙蒂永Aostachatillon.SetParent(f)
	
	f.阿尔纳德Arnad = BArnad阿尔纳德
	f.阿尔纳德Arnad.SetParent(f)
	
	f.沙泰尔阿尔让Chatel_argent = BChatel_argent沙泰尔阿尔让
	f.沙泰尔阿尔让Chatel_argent.SetParent(f)
	
	f.费尼斯Fenis = BFenis费尼斯
	f.费尼斯Fenis.SetParent(f)
	
	f.圣雅凯姆St_jacqueme = BSt_jacqueme圣雅凯姆
	f.圣雅凯姆St_jacqueme.SetParent(f)
	
	f.圣乌尔St_ours = BSt_ours圣乌尔
	f.圣乌尔St_ours.SetParent(f)
	
	f.圣皮埃尔St_pierre = BSt_pierre圣皮埃尔
	f.圣皮埃尔St_pierre.SetParent(f)
	
}
