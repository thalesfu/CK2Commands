package kotthasara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KotthasaraCounty interface {
	feud.County
	BGirikandaka耆厘建荼迦() feud.Barony
	BGokanna瞿羯罗拏() feud.Barony
	BMadakalapuva摩陀迦兰浮瓦() feud.Barony
	BPulatthinagara补罗悉底耶那揭罗() feud.Barony
	BSigiriya斯耆利耶() feud.Barony
	BTrincomalee亭可马里() feud.Barony
	BVatagiri伐多耆厘() feud.Barony
}

type 拘吒娑罗KotthasaraCounty struct {
	feud.BaseCounty
	耆厘建荼迦Girikandaka       feud.Barony
	瞿羯罗拏Gokanna            feud.Barony
	摩陀迦兰浮瓦Madakalapuva     feud.Barony
	补罗悉底耶那揭罗Pulatthinagara feud.Barony
	斯耆利耶Sigiriya           feud.Barony
	亭可马里Trincomalee        feud.Barony
	伐多耆厘Vatagiri           feud.Barony
}

func (c *拘吒娑罗KotthasaraCounty) BGirikandaka耆厘建荼迦() feud.Barony {
	return c.耆厘建荼迦Girikandaka
}

func (c *拘吒娑罗KotthasaraCounty) BGokanna瞿羯罗拏() feud.Barony {
	return c.瞿羯罗拏Gokanna
}

func (c *拘吒娑罗KotthasaraCounty) BMadakalapuva摩陀迦兰浮瓦() feud.Barony {
	return c.摩陀迦兰浮瓦Madakalapuva
}

func (c *拘吒娑罗KotthasaraCounty) BPulatthinagara补罗悉底耶那揭罗() feud.Barony {
	return c.补罗悉底耶那揭罗Pulatthinagara
}

func (c *拘吒娑罗KotthasaraCounty) BSigiriya斯耆利耶() feud.Barony {
	return c.斯耆利耶Sigiriya
}

func (c *拘吒娑罗KotthasaraCounty) BTrincomalee亭可马里() feud.Barony {
	return c.亭可马里Trincomalee
}

func (c *拘吒娑罗KotthasaraCounty) BVatagiri伐多耆厘() feud.Barony {
	return c.伐多耆厘Vatagiri
}

var CKotthasara拘吒娑罗 KotthasaraCounty = &拘吒娑罗KotthasaraCounty{}

func init() {
	f := CKotthasara拘吒娑罗.(*拘吒娑罗KotthasaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1414",
		Title:     "kotthasara",
		TitleName: "拘吒娑罗",
		TitleCode: "c_kotthasara",
		Baronies:  map[string]feud.Barony{},
	}

	f.耆厘建荼迦Girikandaka = BGirikandaka耆厘建荼迦
	f.耆厘建荼迦Girikandaka.SetParent(f)

	f.瞿羯罗拏Gokanna = BGokanna瞿羯罗拏
	f.瞿羯罗拏Gokanna.SetParent(f)

	f.摩陀迦兰浮瓦Madakalapuva = BMadakalapuva摩陀迦兰浮瓦
	f.摩陀迦兰浮瓦Madakalapuva.SetParent(f)

	f.补罗悉底耶那揭罗Pulatthinagara = BPulatthinagara补罗悉底耶那揭罗
	f.补罗悉底耶那揭罗Pulatthinagara.SetParent(f)

	f.斯耆利耶Sigiriya = BSigiriya斯耆利耶
	f.斯耆利耶Sigiriya.SetParent(f)

	f.亭可马里Trincomalee = BTrincomalee亭可马里
	f.亭可马里Trincomalee.SetParent(f)

	f.伐多耆厘Vatagiri = BVatagiri伐多耆厘
	f.伐多耆厘Vatagiri.SetParent(f)

}
