/**
 * @Author Kedaya
 * @Date 2024/12/27 10:07:00
 * @Desc
 **/
package source

// 4G Dongle信息
type DongleInfo struct {
	Imel              string     `json:"imel"`                // dongle imei
	DongleType        int        `json:"dongle_type"`         // Dongle 类型 ("6":"旧 Dongle","10":"支持 eSIM 的新 Dongle")
	Eid               string     `json:"eid"`                 // dongle eid
	EsimActivateState int        `json:"esim_activate_state"` // eSIM 激活状态 ("0":"未激活","1":"已激活")
	SimCardState      int        `json:"sim_card_state"`      // SIM 卡状态 ("0":"未插入","1":"已插入")
	SimSlot           int        `json:"sim_slot"`            // SIM 卡槽使能状态 ("0":"未知","1":"实体 SIM 卡","2":"eSIM")
	EsimInfos         []EsimInfo `json:"esim_infos"`          // eSIM 信息
	SimInfo           SimInfo    `json:"sim_info"`            // SIM 卡信息
}

// SIM 卡信息
type SimInfo struct {
	TelecomOperator int    `json:"telecom_operator"` // 支持的运营商 ("0":"未知","1":"移动","2":"联通","3":"电信")
	SimType         int    `json:"sim_type"`         // SIM 卡类型 ("0":"未知","1":"其他普通 SIM 卡","2":"三网卡")
	Iccid           string `json:"iccid"`            // sim iccid
}

// eSIM 信息
type EsimInfo struct {
	TelecomOperator int    `json:"telecom_operator"` // 支持的运营商 ("0":"未知","1":"移动","2":"联通","3":"电信")
	Enabled         bool   `json:"enabled"`          // eSIM 使能状态 ("false":"未使用","true":"使用中")
	Iccid           string `json:"iccid"`            // sim iccid
}
