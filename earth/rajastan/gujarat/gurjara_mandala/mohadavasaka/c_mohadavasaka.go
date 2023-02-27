package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MohadavasakaCounty interface {
    feud.County
    BDanta檀波() 	feud.Barony
    BIdar伊德尔() 	feud.Barony
    BJhalod旬卢() 	feud.Barony
    BLunawada卢那瓦陀() 	feud.Barony
    BMohadavasaka慕诃陀婆娑迦() 	feud.Barony
    BShamlaji萨罗吉() 	feud.Barony
    BUnjha温恰() 	feud.Barony
}

type 慕诃陀婆娑迦MohadavasakaCounty struct {
	feud.BaseCounty
	檀波Danta 	feud.Barony
	伊德尔Idar 	feud.Barony
	旬卢Jhalod 	feud.Barony
	卢那瓦陀Lunawada 	feud.Barony
	慕诃陀婆娑迦Mohadavasaka 	feud.Barony
	萨罗吉Shamlaji 	feud.Barony
	温恰Unjha 	feud.Barony
}

func (c *慕诃陀婆娑迦MohadavasakaCounty) BDanta檀波() feud.Barony {
	return c.檀波Danta
}
    
func (c *慕诃陀婆娑迦MohadavasakaCounty) BIdar伊德尔() feud.Barony {
	return c.伊德尔Idar
}
    
func (c *慕诃陀婆娑迦MohadavasakaCounty) BJhalod旬卢() feud.Barony {
	return c.旬卢Jhalod
}
    
func (c *慕诃陀婆娑迦MohadavasakaCounty) BLunawada卢那瓦陀() feud.Barony {
	return c.卢那瓦陀Lunawada
}
    
func (c *慕诃陀婆娑迦MohadavasakaCounty) BMohadavasaka慕诃陀婆娑迦() feud.Barony {
	return c.慕诃陀婆娑迦Mohadavasaka
}
    
func (c *慕诃陀婆娑迦MohadavasakaCounty) BShamlaji萨罗吉() feud.Barony {
	return c.萨罗吉Shamlaji
}
    
func (c *慕诃陀婆娑迦MohadavasakaCounty) BUnjha温恰() feud.Barony {
	return c.温恰Unjha
}
    
var CMohadavasaka慕诃陀婆娑迦 MohadavasakaCounty = &慕诃陀婆娑迦MohadavasakaCounty{}

func init() {
	f := CMohadavasaka慕诃陀婆娑迦.(*慕诃陀婆娑迦MohadavasakaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1291",
		Title:     "mohadavasaka",
		TitleName: "慕诃陀婆娑迦",
		TitleCode: "c_mohadavasaka",
		Baronies:  map[string]feud.Barony{},
	}

	f.檀波Danta = BDanta檀波
	f.檀波Danta.SetParent(f)
	
	f.伊德尔Idar = BIdar伊德尔
	f.伊德尔Idar.SetParent(f)
	
	f.旬卢Jhalod = BJhalod旬卢
	f.旬卢Jhalod.SetParent(f)
	
	f.卢那瓦陀Lunawada = BLunawada卢那瓦陀
	f.卢那瓦陀Lunawada.SetParent(f)
	
	f.慕诃陀婆娑迦Mohadavasaka = BMohadavasaka慕诃陀婆娑迦
	f.慕诃陀婆娑迦Mohadavasaka.SetParent(f)
	
	f.萨罗吉Shamlaji = BShamlaji萨罗吉
	f.萨罗吉Shamlaji.SetParent(f)
	
	f.温恰Unjha = BUnjha温恰
	f.温恰Unjha.SetParent(f)
	
}
