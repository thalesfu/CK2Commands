package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HonnoreCounty interface {
    feud.County
    BBalisna婆梨娑那() 	feud.Barony
    BBarkur巴库尔() 	feud.Barony
    BBhatkal婆吒迦罗() 	feud.Barony
    BGokarna瞿羯罗拏() 	feud.Barony
    BHonnore胡奴梨() 	feud.Barony
    BKarwar卡瓦尔() 	feud.Barony
    BMookambika摩甘比加() 	feud.Barony
}

type 胡奴梨HonnoreCounty struct {
	feud.BaseCounty
	婆梨娑那Balisna 	feud.Barony
	巴库尔Barkur 	feud.Barony
	婆吒迦罗Bhatkal 	feud.Barony
	瞿羯罗拏Gokarna 	feud.Barony
	胡奴梨Honnore 	feud.Barony
	卡瓦尔Karwar 	feud.Barony
	摩甘比加Mookambika 	feud.Barony
}

func (c *胡奴梨HonnoreCounty) BBalisna婆梨娑那() feud.Barony {
	return c.婆梨娑那Balisna
}
    
func (c *胡奴梨HonnoreCounty) BBarkur巴库尔() feud.Barony {
	return c.巴库尔Barkur
}
    
func (c *胡奴梨HonnoreCounty) BBhatkal婆吒迦罗() feud.Barony {
	return c.婆吒迦罗Bhatkal
}
    
func (c *胡奴梨HonnoreCounty) BGokarna瞿羯罗拏() feud.Barony {
	return c.瞿羯罗拏Gokarna
}
    
func (c *胡奴梨HonnoreCounty) BHonnore胡奴梨() feud.Barony {
	return c.胡奴梨Honnore
}
    
func (c *胡奴梨HonnoreCounty) BKarwar卡瓦尔() feud.Barony {
	return c.卡瓦尔Karwar
}
    
func (c *胡奴梨HonnoreCounty) BMookambika摩甘比加() feud.Barony {
	return c.摩甘比加Mookambika
}
    
var CHonnore胡奴梨 HonnoreCounty = &胡奴梨HonnoreCounty{}

func init() {
	f := CHonnore胡奴梨.(*胡奴梨HonnoreCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1124",
		Title:     "honnore",
		TitleName: "胡奴梨",
		TitleCode: "c_honnore",
		Baronies:  map[string]feud.Barony{},
	}

	f.婆梨娑那Balisna = BBalisna婆梨娑那
	f.婆梨娑那Balisna.SetParent(f)
	
	f.巴库尔Barkur = BBarkur巴库尔
	f.巴库尔Barkur.SetParent(f)
	
	f.婆吒迦罗Bhatkal = BBhatkal婆吒迦罗
	f.婆吒迦罗Bhatkal.SetParent(f)
	
	f.瞿羯罗拏Gokarna = BGokarna瞿羯罗拏
	f.瞿羯罗拏Gokarna.SetParent(f)
	
	f.胡奴梨Honnore = BHonnore胡奴梨
	f.胡奴梨Honnore.SetParent(f)
	
	f.卡瓦尔Karwar = BKarwar卡瓦尔
	f.卡瓦尔Karwar.SetParent(f)
	
	f.摩甘比加Mookambika = BMookambika摩甘比加
	f.摩甘比加Mookambika.SetParent(f)
	
}
