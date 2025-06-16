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
	// temp := new(expense)
	// temp.fromCsv(data[0])
	// fmt.Println("", temp.toCsv())
	// temp := &expense{0, time.Now(), "Temporary", 20}
	// fmt.Println("", temp.toCsv())
	// writeCsv("data.csv", [][]string{temp.toCsv()})
	// fmt.Println("Param: ", os.Args[1])
}
