package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DjadoCounty interface {
    feud.County
    BAl_qalah盖拉() 	feud.Barony
    BDjado贾多() 	feud.Barony
    BDjemmari杰马拉() 	feud.Barony
    BEl_bedan贝丹() 	feud.Barony
    BGrennah格兰纳() 	feud.Barony
    BMsus姆苏斯() 	feud.Barony
    BTummo图莫() 	feud.Barony
}

type 贾多DjadoCounty struct {
	feud.BaseCounty
	盖拉Al_qalah 	feud.Barony
	贾多Djado 	feud.Barony
	杰马拉Djemmari 	feud.Barony
	贝丹El_bedan 	feud.Barony
	格兰纳Grennah 	feud.Barony
	姆苏斯Msus 	feud.Barony
	图莫Tummo 	feud.Barony
}

func (c *贾多DjadoCounty) BAl_qalah盖拉() feud.Barony {
	return c.盖拉Al_qalah
}
    
func (c *贾多DjadoCounty) BDjado贾多() feud.Barony {
	return c.贾多Djado
}
    
func (c *贾多DjadoCounty) BDjemmari杰马拉() feud.Barony {
	return c.杰马拉Djemmari
}
    
func (c *贾多DjadoCounty) BEl_bedan贝丹() feud.Barony {
	return c.贝丹El_bedan
}
    
func (c *贾多DjadoCounty) BGrennah格兰纳() feud.Barony {
	return c.格兰纳Grennah
}
    
func (c *贾多DjadoCounty) BMsus姆苏斯() feud.Barony {
	return c.姆苏斯Msus
}
    
func (c *贾多DjadoCounty) BTummo图莫() feud.Barony {
	return c.图莫Tummo
}
    
var CDjado贾多 DjadoCounty = &贾多DjadoCounty{}

func init() {
	f := CDjado贾多.(*贾多DjadoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1732",
		Title:     "djado",
		TitleName: "贾多",
		TitleCode: "c_djado",
		Baronies:  map[string]feud.Barony{},
	}

	f.盖拉Al_qalah = BAl_qalah盖拉
	f.盖拉Al_qalah.SetParent(f)
	
	f.贾多Djado = BDjado贾多
	f.贾多Djado.SetParent(f)
	
	f.杰马拉Djemmari = BDjemmari杰马拉
	f.杰马拉Djemmari.SetParent(f)
	
	f.贝丹El_bedan = BEl_bedan贝丹
	f.贝丹El_bedan.SetParent(f)
	
	f.格兰纳Grennah = BGrennah格兰纳
	f.格兰纳Grennah.SetParent(f)
	
	f.姆苏斯Msus = BMsus姆苏斯
	f.姆苏斯Msus.SetParent(f)
	
	f.图莫Tummo = BTummo图莫
	f.图莫Tummo.SetParent(f)
	
}
