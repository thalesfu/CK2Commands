package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YungguanCounty interface {
    feud.County
    BGuihuayuan桂花园() 	feud.Barony
    BNieercun聂耳村() 	feud.Barony
    BPochao坡巢() 	feud.Barony
    BShouchang寿昌() 	feud.Barony
    BTuotou陀头() 	feud.Barony
    BYungguan阳关() 	feud.Barony
    BZiting子亭() 	feud.Barony
}

type 阳关YungguanCounty struct {
	feud.BaseCounty
	桂花园Guihuayuan 	feud.Barony
	聂耳村Nieercun 	feud.Barony
	坡巢Pochao 	feud.Barony
	寿昌Shouchang 	feud.Barony
	陀头Tuotou 	feud.Barony
	阳关Yungguan 	feud.Barony
	子亭Ziting 	feud.Barony
}

func (c *阳关YungguanCounty) BGuihuayuan桂花园() feud.Barony {
	return c.桂花园Guihuayuan
}
    
func (c *阳关YungguanCounty) BNieercun聂耳村() feud.Barony {
	return c.聂耳村Nieercun
}
    
func (c *阳关YungguanCounty) BPochao坡巢() feud.Barony {
	return c.坡巢Pochao
}
    
func (c *阳关YungguanCounty) BShouchang寿昌() feud.Barony {
	return c.寿昌Shouchang
}
    
func (c *阳关YungguanCounty) BTuotou陀头() feud.Barony {
	return c.陀头Tuotou
}
    
func (c *阳关YungguanCounty) BYungguan阳关() feud.Barony {
	return c.阳关Yungguan
}
    
func (c *阳关YungguanCounty) BZiting子亭() feud.Barony {
	return c.子亭Ziting
}
    
var CYungguan阳关 YungguanCounty = &阳关YungguanCounty{}

func init() {
	f := CYungguan阳关.(*阳关YungguanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1515",
		Title:     "yungguan",
		TitleName: "阳关",
		TitleCode: "c_yungguan",
		Baronies:  map[string]feud.Barony{},
	}

	f.桂花园Guihuayuan = BGuihuayuan桂花园
	f.桂花园Guihuayuan.SetParent(f)
	
	f.聂耳村Nieercun = BNieercun聂耳村
	f.聂耳村Nieercun.SetParent(f)
	
	f.坡巢Pochao = BPochao坡巢
	f.坡巢Pochao.SetParent(f)
	
	f.寿昌Shouchang = BShouchang寿昌
	f.寿昌Shouchang.SetParent(f)
	
	f.陀头Tuotou = BTuotou陀头
	f.陀头Tuotou.SetParent(f)
	
	f.阳关Yungguan = BYungguan阳关
	f.阳关Yungguan.SetParent(f)
	
	f.子亭Ziting = BZiting子亭
	f.子亭Ziting.SetParent(f)
	
}
