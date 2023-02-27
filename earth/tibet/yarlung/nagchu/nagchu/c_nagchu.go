package nagchu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NagchuCounty interface {
    feud.County
    BDaqen达前() 	feud.Barony
    BGolug古露() 	feud.Barony
    BKhormang库尔茫() 	feud.Barony
    BLhomar罗玛() 	feud.Barony
    BNagchu那曲() 	feud.Barony
    BNamra那玛切() 	feud.Barony
    BYochak油恰() 	feud.Barony
}

type 那曲NagchuCounty struct {
	feud.BaseCounty
	达前Daqen 	feud.Barony
	古露Golug 	feud.Barony
	库尔茫Khormang 	feud.Barony
	罗玛Lhomar 	feud.Barony
	那曲Nagchu 	feud.Barony
	那玛切Namra 	feud.Barony
	油恰Yochak 	feud.Barony
}

func (c *那曲NagchuCounty) BDaqen达前() feud.Barony {
	return c.达前Daqen
}
    
func (c *那曲NagchuCounty) BGolug古露() feud.Barony {
	return c.古露Golug
}
    
func (c *那曲NagchuCounty) BKhormang库尔茫() feud.Barony {
	return c.库尔茫Khormang
}
    
func (c *那曲NagchuCounty) BLhomar罗玛() feud.Barony {
	return c.罗玛Lhomar
}
    
func (c *那曲NagchuCounty) BNagchu那曲() feud.Barony {
	return c.那曲Nagchu
}
    
func (c *那曲NagchuCounty) BNamra那玛切() feud.Barony {
	return c.那玛切Namra
}
    
func (c *那曲NagchuCounty) BYochak油恰() feud.Barony {
	return c.油恰Yochak
}
    
var CNagchu那曲 NagchuCounty = &那曲NagchuCounty{}

func init() {
	f := CNagchu那曲.(*那曲NagchuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1351",
		Title:     "nagchu",
		TitleName: "那曲",
		TitleCode: "c_nagchu",
		Baronies:  map[string]feud.Barony{},
	}

	f.达前Daqen = BDaqen达前
	f.达前Daqen.SetParent(f)
	
	f.古露Golug = BGolug古露
	f.古露Golug.SetParent(f)
	
	f.库尔茫Khormang = BKhormang库尔茫
	f.库尔茫Khormang.SetParent(f)
	
	f.罗玛Lhomar = BLhomar罗玛
	f.罗玛Lhomar.SetParent(f)
	
	f.那曲Nagchu = BNagchu那曲
	f.那曲Nagchu.SetParent(f)
	
	f.那玛切Namra = BNamra那玛切
	f.那玛切Namra.SetParent(f)
	
	f.油恰Yochak = BYochak油恰
	f.油恰Yochak.SetParent(f)
	
}
