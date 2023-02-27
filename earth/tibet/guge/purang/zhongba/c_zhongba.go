package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZhongbaCounty interface {
    feud.County
    BBarma帕玛() 	feud.Barony
    BGela_b吉拉() 	feud.Barony
    BLabrang拉让() 	feud.Barony
    BLunggar隆格尔() 	feud.Barony
    BParyang帕羊() 	feud.Barony
    BRingtor仁多() 	feud.Barony
    BYagra亚热() 	feud.Barony
}

type 仲巴ZhongbaCounty struct {
	feud.BaseCounty
	帕玛Barma 	feud.Barony
	吉拉Gela_b 	feud.Barony
	拉让Labrang 	feud.Barony
	隆格尔Lunggar 	feud.Barony
	帕羊Paryang 	feud.Barony
	仁多Ringtor 	feud.Barony
	亚热Yagra 	feud.Barony
}

func (c *仲巴ZhongbaCounty) BBarma帕玛() feud.Barony {
	return c.帕玛Barma
}
    
func (c *仲巴ZhongbaCounty) BGela_b吉拉() feud.Barony {
	return c.吉拉Gela_b
}
    
func (c *仲巴ZhongbaCounty) BLabrang拉让() feud.Barony {
	return c.拉让Labrang
}
    
func (c *仲巴ZhongbaCounty) BLunggar隆格尔() feud.Barony {
	return c.隆格尔Lunggar
}
    
func (c *仲巴ZhongbaCounty) BParyang帕羊() feud.Barony {
	return c.帕羊Paryang
}
    
func (c *仲巴ZhongbaCounty) BRingtor仁多() feud.Barony {
	return c.仁多Ringtor
}
    
func (c *仲巴ZhongbaCounty) BYagra亚热() feud.Barony {
	return c.亚热Yagra
}
    
var CZhongba仲巴 ZhongbaCounty = &仲巴ZhongbaCounty{}

func init() {
	f := CZhongba仲巴.(*仲巴ZhongbaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1562",
		Title:     "zhongba",
		TitleName: "仲巴",
		TitleCode: "c_zhongba",
		Baronies:  map[string]feud.Barony{},
	}

	f.帕玛Barma = BBarma帕玛
	f.帕玛Barma.SetParent(f)
	
	f.吉拉Gela_b = BGela_b吉拉
	f.吉拉Gela_b.SetParent(f)
	
	f.拉让Labrang = BLabrang拉让
	f.拉让Labrang.SetParent(f)
	
	f.隆格尔Lunggar = BLunggar隆格尔
	f.隆格尔Lunggar.SetParent(f)
	
	f.帕羊Paryang = BParyang帕羊
	f.帕羊Paryang.SetParent(f)
	
	f.仁多Ringtor = BRingtor仁多
	f.仁多Ringtor.SetParent(f)
	
	f.亚热Yagra = BYagra亚热
	f.亚热Yagra.SetParent(f)
	
}
