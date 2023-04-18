```sql
create table iris (sepallength int, sepalwidth int, petallength int, petalwidth int, species int);
```





```sql
create model irisKnn as knn(sepallength, sepalwidth, petallength, petalwidth, species, 3, euclidean) from iris;
```



```sql
train model iris_knn with euclidean;
```



```sql
predict(iris_knn, [5.1, 3.5, 1.4, 0.2]) from iris_table;
```



处理一个特殊情况，既在predict语句中，例如“PREDICT(iris_knn, [5.1, 3.5, 1.4, 0.2]) FROM iris_table;”， [5.1, 3.5, 1.4, 0.2]会作为一个整体而存在

当语句为create model, train model或者predict的时候，调用tokenize函数

判断语句是否为create model, train model或者predict

