package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"learning/03_ferate/calc"
	//go_tool "go.tools"
)

func main() {
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := &cobra.Command{
		Use:   "healthCheck",
		Short: "体脂计算器",
		Long:  "基于 BMI 的体脂计算器",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name：", name)
			fmt.Println("sex：", sex)
			fmt.Println("tall：", tall)
			fmt.Println("weight：", weight)
			fmt.Println("age：", age)

			// 计算
			bmi := calc.CalcBMI(tall, weight)
			fatRate := calc.CalcFatRate(bmi, age, sex)
			fmt.Println("fatRate:", fatRate)

			// 评估结果
			//fmt.Println(go_tool.Max(3, 5))
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()
}
