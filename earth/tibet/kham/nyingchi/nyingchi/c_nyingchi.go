package nyingchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NyingchiCounty interface {
    feud.County
    BBeba百巴() 	feud.Barony
    BBuchasergyi_lakang布久色吉拉康() 	feud.Barony
    BDrakchi巴宜() 	feud.Barony
    BGongbogyamda工布江达() 	feud.Barony
    BLamaling喇嘛林() 	feud.Barony
    BNyingchi尼池() 	feud.Barony
    BPagsum八松() 	feud.Barony
}

type 尼池NyingchiCounty struct {
	feud.BaseCounty
	百巴Beba 	feud.Barony
	布久色吉拉康Buchasergyi_lakang 	feud.Barony
	巴宜Drakchi 	feud.Barony
	工布江达Gongbogyamda 	feud.Barony
	喇嘛林Lamaling 	feud.Barony
	尼池Nyingchi 	feud.Barony
	八松Pagsum 	feud.Barony
}

func (c *尼池NyingchiCounty) BBeba百巴() feud.Barony {
	return c.百巴Beba
}
    
func (c *尼池NyingchiCounty) BBuchasergyi_lakang布久色吉拉康() feud.Barony {
	return c.布久色吉拉康Buchasergyi_lakang
}
    
func (c *尼池NyingchiCounty) BDrakchi巴宜() feud.Barony {
	return c.巴宜Drakchi
}
    
func (c *尼池NyingchiCounty) BGongbogyamda工布江达() feud.Barony {
	return c.工布江达Gongbogyamda
}
    
func (c *尼池NyingchiCounty) BLamaling喇嘛林() feud.Barony {
	return c.喇嘛林Lamaling
}
    
func (c *尼池NyingchiCounty) BNyingchi尼池() feud.Barony {
	return c.尼池Nyingchi
}
    
func (c *尼池NyingchiCounty) BPagsum八松() feud.Barony {
	return c.八松Pagsum
}
    
var CNyingchi尼池 NyingchiCounty = &尼池NyingchiCounty{}

func init() {
	f := CNyingchi尼池.(*尼池NyingchiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1500",
		Title:     "nyingchi",
		TitleName: "尼池",
		TitleCode: "c_nyingchi",
		Baronies:  map[string]feud.Barony{},
	}

	f.百巴Beba = BBeba百巴
	f.百巴Beba.SetParent(f)
	
	f.布久色吉拉康Buchasergyi_lakang = BBuchasergyi_lakang布久色吉拉康
	f.布久色吉拉康Buchasergyi_lakang.SetParent(f)
	
	f.巴宜Drakchi = BDrakchi巴宜
	f.巴宜Drakchi.SetParent(f)
	
	f.工布江达Gongbogyamda = BGongbogyamda工布江达
	f.工布江达Gongbogyamda.SetParent(f)
	
	f.喇嘛林Lamaling = BLamaling喇嘛林
	f.喇嘛林Lamaling.SetParent(f)
	
	f.尼池Nyingchi = BNyingchi尼池
	f.尼池Nyingchi.SetParent(f)
	
	f.八松Pagsum = BPagsum八松
	f.八松Pagsum.SetParent(f)
	
}
