package calc

import "fmt"

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {
	//bmi, _ = gobmi.BMI(weight, tall)
	if tall <= 0 {
		return 0, fmt.Errorf("身高不能是 0 或者负数")
	}
	// todo 验证体重合法性
	return weight / (tall * tall), nil
}
