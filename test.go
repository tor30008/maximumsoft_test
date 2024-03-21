package main

import "fmt"

type Arr_stuc struct {
	data string
}

func main() {

	var arr_one []Arr_stuc
	var arr_two []Arr_stuc
	var arr Arr_stuc
	i := 0
	for i != -1 {
		fmt.Printf("Enter data Arr_one[%d] (Exit Arr_one Enter 'Break') : ", i)
		fmt.Scan(&arr.data)
		if arr.data != "Break" {
			arr_one = append(arr_one, arr)
			i++
		} else {
			break
		}
	}
	i = 0
	for i != -1 {
		fmt.Printf("Enter data Arr_two[%d] (Exit Arr_two Enter 'Break') : ", i)
		fmt.Scan(&arr.data)
		if arr.data != "Break" {
			arr_two = append(arr_two, arr)
			i++
		} else {
			break
		}
	}
	var result_same []Arr_stuc
	var result_diff []Arr_stuc
	result_same, result_diff = process(arr_one, arr_two)
	fmt.Printf("Result same : %v\n", result_same)
	fmt.Printf("Result diff : %v", result_diff)
}

func process(arr_one []Arr_stuc, arr_two []Arr_stuc) ([]Arr_stuc, []Arr_stuc) {
	fmt.Printf("Result Arr_one : %v \n", arr_one)
	fmt.Printf("Result Arr_two : %v \n", arr_two)

	var same []Arr_stuc
	//var dup []Arr_stuc
	var diff []Arr_stuc
	for _, data_one := range arr_one {
		//fmt.Print(data_one)
		var sta = false
		for _, data_two := range arr_two {
			if data_one == data_two {
				fmt.Printf("Data_one = %v | Data_two = %v \n", data_one, data_two)
				sta = true
				break
			}
		}
		if sta {
			var c_s = true
			for _, s := range same {
				if s == data_one {
					fmt.Printf("s : %v |  data_one : %v \n", s, data_one)
					c_s = false
					break
				}
			}
			if c_s {
				same = append(same, data_one)
			}

		}
	} // หาค่าเหมือนกันก่อน ได้ค่าเหมือนกันแล้ว

	for _, data_one := range arr_one {
		var sta = false
		for _, s := range same {
			if data_one == s {
				sta = true
				break
			}
		}
		if !sta {
			diff = append(diff, data_one)
		}
	}

	for _, data_two := range arr_two {
		var sta = false
		for _, s := range same {
			if data_two == s {
				sta = true
				break
			}
		}
		if !sta {
			diff = append(diff, data_two)
		}
	}

	for _, d := range diff {
		same = append(same, d)
	}

	return same, diff
}

func check_same() {

}
