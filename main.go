package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ===================Global vars===================
var csvName = "data.csv" //Имя файла с данными таблицы

// Объект траты
type expense struct {
	id     int
	date   time.Time
	desc   string
	amount float32
}

// Метод конвертации траты в CSV массив строк
func (e *expense) toCsv() []string {
	return []string{
		strconv.Itoa(e.id),
		e.date.Format("02-01-2006"),
		e.desc,
		fmt.Sprintf("%f", e.amount),
	}
}

// Добавляет текущий объект в файл fileName
func (e *expense) appendToCsv(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(e.toCsv())
	if err != nil {
		panic(err)
	}
}

// Метод конвертации траты из CSV массива строк
func (e *expense) fromCsv(data []string) {
	e.id, _ = strconv.Atoi(data[0])
	e.date, _ = time.Parse("02-01-2006", data[1])
	e.desc = data[2]
	temp, _ := strconv.ParseFloat(data[3], 32)
	e.amount = float32(temp)
}

// Метод вывода в консоль объекта
func (e *expense) print() {
	fmt.Printf("%d\t%s\t%s\t%f\n",
		e.id, e.date.Format("02-01-2006"), e.desc, e.amount)
}

// Читает из файла 'filename'.csv данные
func readCsv(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	ans, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return ans
}

// Записывает в файл fileName массив данных data
func writeCsv(fileName string, data [][]string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	err = writer.WriteAll(data)
	if err != nil {
		panic(err)
	}
}

// Получение текущего id
func getId() int {
	name := "id.txt"
	var id int
	data, err := os.ReadFile(name)
	if err != nil {
		id = 0
	} else {
		temp := string(data)
		id, _ = strconv.Atoi(temp)
		id++
	}
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(strconv.Itoa(id))
	return id
}

// Удалить элемент по id со сдвигом из arr
func removeElem(arr []expense, id int) []expense {
	return append(arr[:id], arr[id+1:]...)
}

func main() {
	// Реализация добавления
	if len(os.Args) == 6 &&
		os.Args[1] == "add" &&
		os.Args[2] == "--desc" &&
		os.Args[4] == "--amount" {
		amount, err := strconv.ParseFloat(os.Args[5], 32)
		if err != nil {
			fmt.Println("Invalid input: amount is not integer")
			return
		}
		exp := expense{getId(), time.Now(), os.Args[3], float32(amount)}
		exp.appendToCsv(csvName)
		fmt.Println("New note added")
		return
	}
	data := readCsv("data.csv")
	var expenseList []expense
	if data != nil {
		for _, elem := range data {
			temp := new(expense)
			temp.fromCsv(elem)
			expenseList = append(expenseList, *temp)
		}
	}
	// Реализация отображения всех элементов
	if len(os.Args) == 2 && os.Args[1] == "list" {
		fmt.Println("ID\tDate\tDescription\tAmount")
		if len(expenseList) == 0 {
			return
		}
		for _, obj := range expenseList {
			obj.print()
		}
		return
		//Реализация получения общей суммы трат
	} else if len(os.Args) == 2 && os.Args[1] == "summary" {
		sum := float32(0.0)
		for _, obj := range expenseList {
			sum += obj.amount
		}
		fmt.Println("Total expenses: ", sum)
		return
		//Реализация удаления траты из таблицы
	} else if len(os.Args) == 4 &&
		os.Args[1] == "delete" &&
		os.Args[2] == "--id" {
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid args: id must be integer")
			return
		}
		for _, obj := range expenseList {
			if obj.id == id {
				expenseList = removeElem(expenseList, id)
				var temp [][]string
				for _, obj := range expenseList {
					temp = append(temp, obj.toCsv())
				}
				writeCsv(csvName, temp)
				fmt.Println("Success delete")
				return
			}
		}
		fmt.Println("Invalid args: can't find row with this id->", id)
		return
	}
	fmt.Println("Invalid args")
}
