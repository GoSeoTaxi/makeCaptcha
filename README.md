# makeCaptcha

WebService to making text captcha + img captcha 


Сервис генерации капчи.

Работает с собственным inMemory хранилищем.
Принцип работы.

В фоне создаётся задача, которая занимается генерацией пары ключ-картинка.
Создаёт их в бесконечном цикле, заполняя буфер. Размер буфера задаётся в количестве единиц, 
через переменную окружения - env:"SIZE_CACHE или через флаг -s=999. Значение по умолчанию 100000
При таком значении потребление памяти порядка 400мб (при необходимости можно завнуть в контейнер или VDS).

Далее из хранилища данные в бесконечном цикле в случайном (на самом деле в псевдослучайном, 
т.е. там не выполняется криптографическая случайность, но в данном случае это не важно) 
порядке попадают в очередь на отдачу. 
Поэтому при большом трафике, неизбежно будет повторение данных (пар ключ-картинка), 
пока генератор не перезапищет новое значение в хранилище 
(с увеличением размера хранилища естественно будет уменьшатся частота повторений).

Данные отдаются по End Point - ip:8082/get
Задать порт и IP интерфейса можно по переменной окружения env:"RUN_ADDRESS или флагом -a
Значение по умолчанию :8082 (все интерфейсы, порт 8082).

Так же есть настройка по времени тайм ауту, который сервис будет пытаться отдать данные (по умолчанию 1с, до ошибки 500).
Увеличить можно через переменную окружения - env:"TIME_OUT или флагом -t

Запуск сервиса, конечно лучше производить как сервис. В сервисе предуспотрено 2 варианта логирования.

1) Debug - включён по умолчанию
2) info - включается выключением Debug режима через env:"CAP_MAKER_APP_SERVER_DEBUG или флаг debug=false

+++++++++++++++++++++++++++++
Производительность:
На выделенном сервере - AMD Ryzen 5 3600 6-Core Processor
Полный процесс перегенерации всего хранилища на 100000 элементов занимает менее 30 минут
При этом RPS составляет более 45.000
99% запросов были обработаны в менее чем 343.00us

