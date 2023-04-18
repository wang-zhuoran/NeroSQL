# A distributed SQL Engine supporting AI model

- Final Year Project



## Features

- Supporting basic SQL clauses, such as `SELECT`, `INSERT`, `DELET`, `CREAT TABLE`, etc.
- Supporting direct calling of Clustering operation within the DB;
- Distributed Architecture: Client-Server Model

## AI-oriented declarative language syntax definition
```sql
-- Declare a model
CREATE MODEL model_name AS
SELECT model(args*) AS pred_output
	...
FROM table;

-- Train the model
TRAIN MODEL model_name
WITH (Args*);

-- Make prediction
SELECT PREDICT(model_name, input) FROM new_table;

```
Example 1:
```sql
-- 定义线性回归模型
CREATE MODEL model_name AS
SELECT linear_regression(col1, col2) AS output
FROM dataset_name;

-- 训练线性回归模型
TRAIN MODEL model_name
WITH optimizer='gradient_descent', learning_rate=0.1, epochs=100;

-- 预测新数据
SELECT predict(model_name, [col1, col2]) FROM newdata_name;
```
Example 2:
```sql
-- 声明模型
CREATE MODEL iris_knn AS
SELECT kNN('features', 'label', 3) AS prediction
FROM iris_table;

-- 训练模型
TRAIN MODEL iris_knn
WITH distance_function='euclidean';

-- 预测新数据
SELECT PREDICT(iris_knn, ARRAY[5.1, 3.5, 1.4, 0.2]), PREDICT(iris_knn, ARRAY[6.9, 3.1, 5.1, 2.3])
FROM iris_table;
```



## To-Do List

- [x] SQL Parser
- [ ] AI operation parser

- [ ] Planner

- [ ] Executor

- [ ] Optimization

- [x] Distributed Architecture 



