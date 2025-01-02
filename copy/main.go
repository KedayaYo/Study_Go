package main

import (
	"fmt"
	"go_learn/copy/source"
	"go_learn/copy/target"
	"reflect"
)

// DeepCopy 使用反射实现深拷贝
func DeepCopy(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// 检查 src 和 dst 是否有效
	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("深拷贝 srcKind和dstKind都必须是指针")
	}

	// 解引用指针
	srcVal = srcVal.Elem()
	dstVal = dstVal.Elem()

	// 检查类型兼容性
	if !srcVal.IsValid() || !dstVal.IsValid() {
		return fmt.Errorf("深拷贝 源或目标值无效")
	}

	if srcVal.Type() != dstVal.Type() {
		return fmt.Errorf("深拷贝 类型不匹配: %s and %s", srcVal.Type(), dstVal.Type())
	}

	// 递归拷贝
	return deepCopyRecursive(srcVal, dstVal)
}

// deepCopyRecursive 递归拷贝
func deepCopyRecursive(srcVal, dstVal reflect.Value) error {
	switch srcVal.Kind() {
	case reflect.Struct:
		for i := 0; i < srcVal.NumField(); i++ {
			if err := deepCopyRecursive(srcVal.Field(i), dstVal.Field(i)); err != nil {
				return err
			}
		}
	case reflect.Slice:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.MakeSlice(srcVal.Type(), srcVal.Len(), srcVal.Cap()))
			for i := 0; i < srcVal.Len(); i++ {
				if err := deepCopyRecursive(srcVal.Index(i), dstVal.Index(i)); err != nil {
					return err
				}
			}
		}
	case reflect.Ptr:
		if !srcVal.IsNil() {
			dstVal.Set(reflect.New(srcVal.Elem().Type()))
			return deepCopyRecursive(srcVal.Elem(), dstVal.Elem())
		}
	default:
		dstVal.Set(srcVal)
	}
	return nil
}

// 转换函数：从 source.DongleInfo 转换为 target.DongleInfo
func convertDongleInfo(djiDongle source.DongleInfo) target.DongleInfo {
	// 转换 eSIM 信息
	esimInfos := make([]target.EsimInfo, len(djiDongle.EsimInfos))
	for i, esim := range djiDongle.EsimInfos {
		esimInfos[i] = target.EsimInfo{
			TelecomOperator: esim.TelecomOperator,
			Enabled:         esim.Enabled,
			Iccid:           esim.Iccid,
		}
	}

	// 转换 SIM 信息
	simInfo := target.SimInfo{
		TelecomOperator: djiDongle.SimInfo.TelecomOperator,
		SimType:         djiDongle.SimInfo.SimType,
		Iccid:           djiDongle.SimInfo.Iccid,
	}

	// 构建最终结果
	return target.DongleInfo{
		Imel:              djiDongle.Imel,
		DongleType:        djiDongle.DongleType,
		Eid:               djiDongle.Eid,
		EsimActivateState: djiDongle.EsimActivateState,
		SimCardState:      djiDongle.SimCardState,
		SimSlot:           djiDongle.SimSlot,
		EsimInfos:         esimInfos,
		SimInfo:           simInfo,
	}
}

// 转换函数：从 []source.DongleInfo 转换为 []target.DongleInfo
func convertDongleInfos(djiDongles []source.DongleInfo) []target.DongleInfo {
	entityDongles := make([]target.DongleInfo, len(djiDongles))
	for i, djiDongle := range djiDongles {
		entityDongles[i] = convertDongleInfo(djiDongle)
	}
	return entityDongles
}

func main() {
	// 示例数据
	djiDongles := []source.DongleInfo{
		{
			Imel:              "123456789012345",
			DongleType:        10,
			Eid:               "EID123456789",
			EsimActivateState: 1,
			SimCardState:      1,
			SimSlot:           2,
			EsimInfos: []source.EsimInfo{
				{TelecomOperator: 1, Enabled: true, Iccid: "ICCID123"},
			},
			SimInfo: source.SimInfo{TelecomOperator: 2, SimType: 1, Iccid: "SIMICCID123"},
		},
	}

	// 转换
	entityDongles := convertDongleInfos(djiDongles)

	// 输出结果
	fmt.Printf("Source (DJI): %+v\n", djiDongles)
	fmt.Printf("Destination (Entity): %+v\n", entityDongles)
}
