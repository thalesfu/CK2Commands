package halaban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HalabanCounty interface {
	feud.County
	BAfif阿菲夫() feud.Barony
	BAlbaharah巴哈雷() feud.Barony
	BAlbijadiyah比贾迪亚() feud.Barony
	BAljammaniyah贾曼尼亚() feud.Barony
	BAlmaklah迈克拉() feud.Barony
	BAlqaiyah加伊耶() feud.Barony
	BArradum雷杜木() feud.Barony
	BBadaiidyan巴达鲁杜安() feud.Barony
}

type 哈莱班HalabanCounty struct {
	feud.BaseCounty
	阿菲夫Afif          feud.Barony
	巴哈雷Albaharah     feud.Barony
	比贾迪亚Albijadiyah  feud.Barony
	贾曼尼亚Aljammaniyah feud.Barony
	迈克拉Almaklah      feud.Barony
	加伊耶Alqaiyah      feud.Barony
	雷杜木Arradum       feud.Barony
	巴达鲁杜安Badaiidyan  feud.Barony
}

func (c *哈莱班HalabanCounty) BAfif阿菲夫() feud.Barony {
	return c.阿菲夫Afif
}

func (c *哈莱班HalabanCounty) BAlbaharah巴哈雷() feud.Barony {
	return c.巴哈雷Albaharah
}

func (c *哈莱班HalabanCounty) BAlbijadiyah比贾迪亚() feud.Barony {
	return c.比贾迪亚Albijadiyah
}

func (c *哈莱班HalabanCounty) BAljammaniyah贾曼尼亚() feud.Barony {
	return c.贾曼尼亚Aljammaniyah
}

func (c *哈莱班HalabanCounty) BAlmaklah迈克拉() feud.Barony {
	return c.迈克拉Almaklah
}

func (c *哈莱班HalabanCounty) BAlqaiyah加伊耶() feud.Barony {
	return c.加伊耶Alqaiyah
}

func (c *哈莱班HalabanCounty) BArradum雷杜木() feud.Barony {
	return c.雷杜木Arradum
}

func (c *哈莱班HalabanCounty) BBadaiidyan巴达鲁杜安() feud.Barony {
	return c.巴达鲁杜安Badaiidyan
}

var CHalaban哈莱班 HalabanCounty = &哈莱班HalabanCounty{}

func init() {
	f := CHalaban哈莱班.(*哈莱班HalabanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "862",
		Title:     "halaban",
		TitleName: "哈莱班",
		TitleCode: "c_halaban",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿菲夫Afif = BAfif阿菲夫
	f.阿菲夫Afif.SetParent(f)

	f.巴哈雷Albaharah = BAlbaharah巴哈雷
	f.巴哈雷Albaharah.SetParent(f)

	f.比贾迪亚Albijadiyah = BAlbijadiyah比贾迪亚
	f.比贾迪亚Albijadiyah.SetParent(f)

	f.贾曼尼亚Aljammaniyah = BAljammaniyah贾曼尼亚
	f.贾曼尼亚Aljammaniyah.SetParent(f)

	f.迈克拉Almaklah = BAlmaklah迈克拉
	f.迈克拉Almaklah.SetParent(f)

	f.加伊耶Alqaiyah = BAlqaiyah加伊耶
	f.加伊耶Alqaiyah.SetParent(f)

	f.雷杜木Arradum = BArradum雷杜木
	f.雷杜木Arradum.SetParent(f)

	f.巴达鲁杜安Badaiidyan = BBadaiidyan巴达鲁杜安
	f.巴达鲁杜安Badaiidyan.SetParent(f)

}
