package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UdayagiriCounty interface {
    feud.County
    BGandikota甘地科塔() 	feud.Barony
    BKanipakam贾尼帕卡姆() 	feud.Barony
    BMahanandi摩诃难提() 	feud.Barony
    BNandyal纳恩迪亚尔() 	feud.Barony
    BSriparvata室利钵伐多() 	feud.Barony
    BSrisailam室利势罗() 	feud.Barony
    BUdayagiri优陀耶耆厘() 	feud.Barony
    BVallurapura伐露罗补罗() 	feud.Barony
}

type 优陀耶耆厘UdayagiriCounty struct {
	feud.BaseCounty
	甘地科塔Gandikota 	feud.Barony
	贾尼帕卡姆Kanipakam 	feud.Barony
	摩诃难提Mahanandi 	feud.Barony
	纳恩迪亚尔Nandyal 	feud.Barony
	室利钵伐多Sriparvata 	feud.Barony
	室利势罗Srisailam 	feud.Barony
	优陀耶耆厘Udayagiri 	feud.Barony
	伐露罗补罗Vallurapura 	feud.Barony
}

func (c *优陀耶耆厘UdayagiriCounty) BGandikota甘地科塔() feud.Barony {
	return c.甘地科塔Gandikota
}
    
func (c *优陀耶耆厘UdayagiriCounty) BKanipakam贾尼帕卡姆() feud.Barony {
	return c.贾尼帕卡姆Kanipakam
}
    
func (c *优陀耶耆厘UdayagiriCounty) BMahanandi摩诃难提() feud.Barony {
	return c.摩诃难提Mahanandi
}
    
func (c *优陀耶耆厘UdayagiriCounty) BNandyal纳恩迪亚尔() feud.Barony {
	return c.纳恩迪亚尔Nandyal
}
    
func (c *优陀耶耆厘UdayagiriCounty) BSriparvata室利钵伐多() feud.Barony {
	return c.室利钵伐多Sriparvata
}
    
func (c *优陀耶耆厘UdayagiriCounty) BSrisailam室利势罗() feud.Barony {
	return c.室利势罗Srisailam
}
    
func (c *优陀耶耆厘UdayagiriCounty) BUdayagiri优陀耶耆厘() feud.Barony {
	return c.优陀耶耆厘Udayagiri
}
    
func (c *优陀耶耆厘UdayagiriCounty) BVallurapura伐露罗补罗() feud.Barony {
	return c.伐露罗补罗Vallurapura
}
    
var CUdayagiri优陀耶耆厘 UdayagiriCounty = &优陀耶耆厘UdayagiriCounty{}

func init() {
	f := CUdayagiri优陀耶耆厘.(*优陀耶耆厘UdayagiriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1122",
		Title:     "udayagiri",
		TitleName: "优陀耶耆厘",
		TitleCode: "c_udayagiri",
		Baronies:  map[string]feud.Barony{},
	}

	f.甘地科塔Gandikota = BGandikota甘地科塔
	f.甘地科塔Gandikota.SetParent(f)
	
	f.贾尼帕卡姆Kanipakam = BKanipakam贾尼帕卡姆
	f.贾尼帕卡姆Kanipakam.SetParent(f)
	
	f.摩诃难提Mahanandi = BMahanandi摩诃难提
	f.摩诃难提Mahanandi.SetParent(f)
	
	f.纳恩迪亚尔Nandyal = BNandyal纳恩迪亚尔
	f.纳恩迪亚尔Nandyal.SetParent(f)
	
	f.室利钵伐多Sriparvata = BSriparvata室利钵伐多
	f.室利钵伐多Sriparvata.SetParent(f)
	
	f.室利势罗Srisailam = BSrisailam室利势罗
	f.室利势罗Srisailam.SetParent(f)
	
	f.优陀耶耆厘Udayagiri = BUdayagiri优陀耶耆厘
	f.优陀耶耆厘Udayagiri.SetParent(f)
	
	f.伐露罗补罗Vallurapura = BVallurapura伐露罗补罗
	f.伐露罗补罗Vallurapura.SetParent(f)
	
}
