SELECT surname
FROM table_name
ORDER BY (surname = 'Иванов') DESC , surname ASC;

-----------------------------------------------


SELECT surname
FROM (
         SELECT surname, 1 AS sort_order
         FROM table_name
         WHERE surname = 'Иванов'
         UNION ALL
         SELECT surname, 2 AS sort_order
         FROM table_name
     ) AS sorted_surnames
ORDER BY sort_order, surname ASC;
