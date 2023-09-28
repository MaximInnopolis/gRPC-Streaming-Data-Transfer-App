# SQL Sorting Challenge

This task involves writing an SQL query to sort and retrieve surnames from a table. The table contains a single field with text data representing last names. The goal is to prioritize "Иванов" surnames first and then sort all other surnames in alphabetical order.

## SQL Query

You can use the following SQL query to achieve this:

```sql
SELECT surname
FROM table_name
ORDER BY (surname = 'Иванов') DESC, surname ASC;
```

In this query:

1. I select the "surname" column from the table.
2. I use the ORDER BY clause to specify the sorting criteria.
3. The first criterion, (surname = 'Иванов') DESC, sorts all "Иванов" surnames first in descending order (1st priority).
4. The second criterion, surname ASC, sorts the remaining surnames in ascending alphabetical order (2nd priority).
5. This query will display all "Иванов" surnames at the beginning of the result set, followed by other surnames sorted in alphabetical order.

SQL Migration Files
To apply this SQL query to your database, you can use SQL migration files. Below are the recommended migration files with SQL code to create and drop the table:

init_mg.up.sql
This SQL file should contain the code to create the table and insert sample data if needed: