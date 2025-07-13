# **Expense tracker CLI**

Простейший консольный менеджер трат на Go.  
Сохраняет траты в таблицу *.csv* файла.  

Столбцы таблицы:
ID | Date | Description | Amount  

Список команд:
- expense-tracker add --desc "описание траты" --amount *стоимость* -> добавление в таблицу новой траты
- expense-tracker list -> вывод таблицы с тратами
- expense-tracker summary -> общая сумма трат
- expense-tracker delete --id *id записи* -> удаление записи по id