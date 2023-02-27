package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CharkliqCounty interface {
    feud.County
    BCharkliq卡克里克() 	feud.Barony
    BKargan库尔干() 	feud.Barony
    BLop罗卜() 	feud.Barony
    BMaojiaxiang毛家乡() 	feud.Barony
    BMerdek麦德克() 	feud.Barony
    BMiran米兰() 	feud.Barony
    BZhehor朱倭() 	feud.Barony
}

type 卡克里克CharkliqCounty struct {
	feud.BaseCounty
	卡克里克Charkliq 	feud.Barony
	库尔干Kargan 	feud.Barony
	罗卜Lop 	feud.Barony
	毛家乡Maojiaxiang 	feud.Barony
	麦德克Merdek 	feud.Barony
	米兰Miran 	feud.Barony
	朱倭Zhehor 	feud.Barony
}

func (c *卡克里克CharkliqCounty) BCharkliq卡克里克() feud.Barony {
	return c.卡克里克Charkliq
}
    
func (c *卡克里克CharkliqCounty) BKargan库尔干() feud.Barony {
	return c.库尔干Kargan
}
    
func (c *卡克里克CharkliqCounty) BLop罗卜() feud.Barony {
	return c.罗卜Lop
}
    
func (c *卡克里克CharkliqCounty) BMaojiaxiang毛家乡() feud.Barony {
	return c.毛家乡Maojiaxiang
}
    
func (c *卡克里克CharkliqCounty) BMerdek麦德克() feud.Barony {
	return c.麦德克Merdek
}
    
func (c *卡克里克CharkliqCounty) BMiran米兰() feud.Barony {
	return c.米兰Miran
}
    
func (c *卡克里克CharkliqCounty) BZhehor朱倭() feud.Barony {
	return c.朱倭Zhehor
}
    
var CCharkliq卡克里克 CharkliqCounty = &卡克里克CharkliqCounty{}

func init() {
	f := CCharkliq卡克里克.(*卡克里克CharkliqCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1442",
		Title:     "charkliq",
		TitleName: "卡克里克",
		TitleCode: "c_charkliq",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡克里克Charkliq = BCharkliq卡克里克
	f.卡克里克Charkliq.SetParent(f)
	
	f.库尔干Kargan = BKargan库尔干
	f.库尔干Kargan.SetParent(f)
	
	f.罗卜Lop = BLop罗卜
	f.罗卜Lop.SetParent(f)
	
	f.毛家乡Maojiaxiang = BMaojiaxiang毛家乡
	f.毛家乡Maojiaxiang.SetParent(f)
	
	f.麦德克Merdek = BMerdek麦德克
	f.麦德克Merdek.SetParent(f)
	
	f.米兰Miran = BMiran米兰
	f.米兰Miran.SetParent(f)
	
	f.朱倭Zhehor = BZhehor朱倭
	f.朱倭Zhehor.SetParent(f)
	
}
