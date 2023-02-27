package krakowskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KrakowskieCounty interface {
    feud.County
    BBochnia博赫尼亚() 	feud.Barony
    BCzestochowa琴斯托霍瓦() 	feud.Barony
    BJedrzejow延杰尤夫() 	feud.Barony
    BKrakow克拉科夫() 	feud.Barony
    BOlsztyn奥尔什丁() 	feud.Barony
    BWieliczka维利奇卡() 	feud.Barony
    BZakopane扎科帕内() 	feud.Barony
}

type 克拉科夫KrakowskieCounty struct {
	feud.BaseCounty
	博赫尼亚Bochnia 	feud.Barony
	琴斯托霍瓦Czestochowa 	feud.Barony
	延杰尤夫Jedrzejow 	feud.Barony
	克拉科夫Krakow 	feud.Barony
	奥尔什丁Olsztyn 	feud.Barony
	维利奇卡Wieliczka 	feud.Barony
	扎科帕内Zakopane 	feud.Barony
}

func (c *克拉科夫KrakowskieCounty) BBochnia博赫尼亚() feud.Barony {
	return c.博赫尼亚Bochnia
}
    
func (c *克拉科夫KrakowskieCounty) BCzestochowa琴斯托霍瓦() feud.Barony {
	return c.琴斯托霍瓦Czestochowa
}
    
func (c *克拉科夫KrakowskieCounty) BJedrzejow延杰尤夫() feud.Barony {
	return c.延杰尤夫Jedrzejow
}
    
func (c *克拉科夫KrakowskieCounty) BKrakow克拉科夫() feud.Barony {
	return c.克拉科夫Krakow
}
    
func (c *克拉科夫KrakowskieCounty) BOlsztyn奥尔什丁() feud.Barony {
	return c.奥尔什丁Olsztyn
}
    
func (c *克拉科夫KrakowskieCounty) BWieliczka维利奇卡() feud.Barony {
	return c.维利奇卡Wieliczka
}
    
func (c *克拉科夫KrakowskieCounty) BZakopane扎科帕内() feud.Barony {
	return c.扎科帕内Zakopane
}
    
var CKrakowskie克拉科夫 KrakowskieCounty = &克拉科夫KrakowskieCounty{}

func init() {
	f := CKrakowskie克拉科夫.(*克拉科夫KrakowskieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "527",
		Title:     "krakowskie",
		TitleName: "克拉科夫",
		TitleCode: "c_krakowskie",
		Baronies:  map[string]feud.Barony{},
	}

	f.博赫尼亚Bochnia = BBochnia博赫尼亚
	f.博赫尼亚Bochnia.SetParent(f)
	
	f.琴斯托霍瓦Czestochowa = BCzestochowa琴斯托霍瓦
	f.琴斯托霍瓦Czestochowa.SetParent(f)
	
	f.延杰尤夫Jedrzejow = BJedrzejow延杰尤夫
	f.延杰尤夫Jedrzejow.SetParent(f)
	
	f.克拉科夫Krakow = BKrakow克拉科夫
	f.克拉科夫Krakow.SetParent(f)
	
	f.奥尔什丁Olsztyn = BOlsztyn奥尔什丁
	f.奥尔什丁Olsztyn.SetParent(f)
	
	f.维利奇卡Wieliczka = BWieliczka维利奇卡
	f.维利奇卡Wieliczka.SetParent(f)
	
	f.扎科帕内Zakopane = BZakopane扎科帕内
	f.扎科帕内Zakopane.SetParent(f)
	
}
