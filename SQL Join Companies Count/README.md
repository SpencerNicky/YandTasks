## Task:

В реляционной базе данных существуют таблицы:

Cities - список городов

id - первичный ключ
name - название
population - численность населения
founded - год основания
country_id - id страны

Countries - список стран

id - первичный ключ
name - название
population - численность населения
gdp - валовый продукт в долларах

Companies - компании

id - первичный ключ
name - название
city_id - город в котором находится штаб-квартира
revenue - годовая выручка в долларах
labors - численность сотрудников

Постройте таблицу, где для каждой страны посчитано число компаний, удовлетворяющих условиям:

1) штаб квартира компании находится в этой стране
2) число сотрудников компании не менее 1000 человек

---

## Solution:

Для построения таблицы необходимо использовать SQL запрос с использованием оператора JOIN, чтобы объединить таблицы Cities, Countries и Companies.
```
SELECT 
  Countries.name AS country_name,
  COUNT(Companies.id) AS company_count
FROM 
  Companies
  JOIN Cities ON Companies.city_id = Cities.id
  JOIN Countries ON Cities.country_id = Countries.id
WHERE 
  Companies.labors >= 1000
GROUP BY 
  Countries.name
```