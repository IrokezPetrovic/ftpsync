## Синхронизатор локальных папок с FTP ##

### Использование ###
```bash
ftpsync /path/to/configfile
```
Если путь к конфигу не указан - ищет config.json в текущей папке

### Пример конфига ###
```js
{
    "profiles":[        
        {
            "server":"monitor.digisky.ru:2124", //Адрес сервера и порт            
            "username":"Administrator@digisky.lan", 
            "password":"Copoakbo123",
            "path":"/Pool0/production/test", //Путь на FTP для этого профиля
            "tasks":[ //Список задач для профиля
                {
                    "from":"/Volumes/WORK/etc", //Откуда (Абсолютный путь)
                    "to":"etc"  //Куда (подпапка в папке "path")
                },
                {
                    "from":"/Volumes/WORK/models",
                    "to":"models",
                    "exclude":[".gif$"]//массив с регулярками для исключения локаьных файлов по пути
                }
            ]
        }
    ]
}
```
Про регуляки для исключения файлов можно посмотреть тут https://github.com/google/re2/wiki/Syntax

Профилей в списке может быть несколько, задач в профиле - тоже можетбыть несколько. 

Все задачи во всех профилях выполняются параллельно
